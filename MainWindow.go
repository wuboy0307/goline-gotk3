package main

import (
	"fmt"
	"github.com/carylorrk/goline-gotk3/api"
	prot "github.com/carylorrk/goline-gotk3/protocol"
	"github.com/carylorrk/goline-gotk3/res/glade"
	"github.com/carylorrk/goline-gotk3/res/image"
	"github.com/carylorrk/gotk3/glib"
	"github.com/carylorrk/gotk3/gtk"
	"sync"
	"time"
)

type MainWindow struct {
	*LineWindow
	friendsBox *gtk.Box
	config     *gtk.Button
	logout     *gtk.Button
	exit       *gtk.Button

	chatWindows map[string]*ChatWindow
}

func NewMainWindow() *MainWindow {
	window := &MainWindow{
		LineWindow:  NewLineWindow(glade.MainWindow),
		chatWindows: make(map[string]*ChatWindow)}
	window.setupWidgets()
	window.setupDefault()
	go window.poll()
	return window
}

func (self *MainWindow) setupWidgets() {
	self.setupFriendsWidgets()
	self.setupMoreWidgets()
}

func (self *MainWindow) setupDefault() {
	self.setupFriendsDefault()
	self.setupMoreDefault()
}

func (self *MainWindow) setupFriendsWidgets() {
	self.friendsBox = self.getIObjectWithType("FriendsBox", &gtk.Box{}).(*gtk.Box)
}

func (self *MainWindow) setupFriendsDefault() {
	self.friendsBox.Add(NewFriendsTypeLabel("Groups").label)
	for _, group := range goline.client.Groups {
		entity := api.NewLineGroupWrapper(group)
		self.attachFriend(entity)
	}
	self.friendsBox.Add(NewFriendsTypeLabel("Rooms").label)
	for _, room := range goline.client.Rooms {
		entity := api.NewLineRoomWrapper(room)
		self.attachFriend(entity)
	}
	self.friendsBox.Add(NewFriendsTypeLabel("Contacts").label)
	for _, contact := range goline.client.Contacts {
		entity := api.NewLineContactWrapper(contact)
		self.attachFriend(entity)
	}
}

func (self *MainWindow) attachFriend(entity api.LineEntity) {
	btn, err := gtk.ButtonNewWithLabel(entity.GetName())
	if err != nil {
		goline.LoggerPanicln(err)
	}
	self.Connect(btn, "clicked", func() {
		id := entity.GetId()
		if self.chatWindows[id] == nil {
			self.chatWindows[id] = NewChatWindow(self, entity)
			self.chatWindows[id].window.ShowAll()
		} else {
			self.chatWindows[id].window.Present()
		}
	})
	self.friendsBox.Add(btn)
}

func (self *MainWindow) setupMoreWidgets() {
	self.config = self.getIObjectWithType("ConfigButton", &gtk.Button{}).(*gtk.Button)
	self.logout = self.getIObjectWithType("LogoutButton", &gtk.Button{}).(*gtk.Button)
	self.exit = self.getIObjectWithType("ExitButton", &gtk.Button{}).(*gtk.Button)
}

func (self *MainWindow) setupMoreDefault() {
	self.Connect(self.window, "destroy", gtk.MainQuit)
	self.Connect(self.config, "clicked", func() {
		NewConfigWindow().window.ShowAll()
	})
	self.Connect(self.logout, "clicked", func() {
		goline.AuthToken = ""
		err := goline.SaveSettings()
		if err != nil {
			goline.LoggerPrintln(err)
			RunAlertMessage(self.window, "Failed to clean authorization token.")
		}
		self.window.Destroy()
	})
	self.Connect(self.exit, "clicked", func() { self.window.Destroy() })
}

func (self *MainWindow) parseHttpResponse(str string) (code int, err error) {
	_, err = fmt.Sscanf(str, "HTTP Response code: %d", &code)
	return
}

var pollLock sync.Mutex

func (self *MainWindow) poll() {
	var retryCount int
	var opRevision int64
	for {
		pollLock.Lock()
		glib.IdleAdd(func() {
			pollLock.Lock()
			defer pollLock.Unlock()
			var maxRevision int64
			var parseErr, clientErr error
			operations, fetchErr := goline.client.FetchNewOperations(50)
			if fetchErr != nil {
				var code int
				code, parseErr = self.parseHttpResponse(fetchErr.Error())
				if parseErr == nil && code == 400 {
					goline.client, clientErr = api.ReloginLineClient(goline.client)
					if clientErr != nil {
						goline.LoggerPrintln(clientErr)
						goto errPath
					}
				} else if parseErr != nil || code < 200 || code >= 300 {
					goline.LoggerPrintln(
						"FetchNewOperations:", fetchErr,
						"parseHttpResponse:", parseErr)
					goto retry
				}
			}
			maxRevision = 0
			for _, operation := range operations {
				revision := operation.GetRevision()
				if revision <= opRevision {
					return
				}

				if revision > maxRevision {
					maxRevision = revision
				}

				opType := operation.GetTypeA1()
				switch opType {
				case prot.OpType_SEND_MESSAGE:
					fallthrough
				case prot.OpType_SEND_CONTENT:
					fallthrough
				case prot.OpType_RECEIVE_MESSAGE:
					message := operation.GetMessage()
					if opType == prot.OpType_SEND_MESSAGE &&
						(message.ContentType == prot.ContentType_VIDEO ||
							message.ContentType == prot.ContentType_IMAGE) {
						continue
					}

					if goline.client.Profile == nil {
						var err error
						goline.client, err = api.ReloginLineClient(goline.client)
						if err != nil {
							goline.LoggerPrintln(err)
							goto errPath
						}
					}

					mid := goline.client.Profile.GetMid()
					fromId := message.GetFrom()
					toId := message.GetTo()
					var id string
					if fromId == mid {
						id = toId
					} else {
						if toId == mid {
							id = fromId
						} else {
							id = toId
						}
					}
					chatWindow := self.chatWindows[id]
					if chatWindow == nil {
						entity, err := goline.client.GetLineEntityById(id)
						if err != nil || entity == nil {
							RunAlertMessage(self.window, "Cannot get message's sender.")
						} else {
							self.chatWindows[id] = NewChatWindow(self, entity)
							self.chatWindows[id].window.ShowAll()
						}
					} else {
						chatWindow.addMessage(message)
						if !chatWindow.window.IsActive() {
							chatWindow.setupIcon(image.IconRed)
						}
					}
				}
			}
			opRevision = maxRevision
			retryCount = 0
			return
		retry:
			retryCount++
			if retryCount <= 10 {
				return
			}
		errPath:
			goline.LoggerPanicln("Failed to get new message! Program closed.")

		})
		pollLock.Unlock()
		time.Sleep(500 * time.Millisecond)
	}
}

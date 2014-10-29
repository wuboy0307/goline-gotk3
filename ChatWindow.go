package main

import (
	"github.com/carylorrk/goline-gotk3/api"
	prot "github.com/carylorrk/goline-gotk3/protocol"
	"github.com/carylorrk/goline-gotk3/res/glade"
	"github.com/carylorrk/goline-gotk3/res/image"
	"github.com/carylorrk/gotk3/gdk"
	"github.com/carylorrk/gotk3/gtk"
)

type ChatWindow struct {
	*LineWindow
	scrolled     *gtk.ScrolledWindow
	conversation *gtk.Box
	input        *gtk.Entry
	send         *gtk.Button

	parent     *MainWindow
	entity     api.LineEntity
	messageBox *prot.TMessageBox

	isAdd bool
}

func NewChatWindow(parent *MainWindow, entity api.LineEntity) *ChatWindow {
	window := &ChatWindow{
		LineWindow: NewLineWindow(glade.ChatWindow),
		parent:     parent,
		entity:     entity}

	window.setupWidgets()
	window.setupMessageBox()
	window.setupDefault()
	return window
}

func (self *ChatWindow) setupWidgets() {
	self.scrolled = self.getIObjectWithType("ScrolledWindow", &gtk.ScrolledWindow{}).(*gtk.ScrolledWindow)
	self.conversation = self.getIObjectWithType("Conversation", &gtk.Box{}).(*gtk.Box)
	self.input = self.getIObjectWithType("Input", &gtk.Entry{}).(*gtk.Entry)
	self.send = self.getIObjectWithType("Send", &gtk.Button{}).(*gtk.Button)

}

func (self *ChatWindow) setupDefault() {
	self.window.SetTitle(self.entity.GetName())
	self.Connect(self.window, "destroy", func() {
		self.parent.chatWindows[self.entity.GetId()] = nil
	})

	self.Connect(self.window, "focus-in-event", func() {
		self.setupIcon(image.Icon)
	})

	self.Connect(self.conversation, "size-allocate", func() {
		if self.isAdd {
			adj := self.scrolled.GetVAdjustment()
			adj.SetValue(adj.GetUpper() - adj.GetPageSize())
			self.isAdd = false
		}
	})

	self.Connect(self.window, "key-press-event", func(window *gtk.Window, ev *gdk.Event) {
		keyEvent := &gdk.EventKey{ev}
		keyVal := keyEvent.KeyVal()
		if keyVal == KEY_Return || keyVal == KEY_KP_Enter {
			self.send.Emit("clicked")
		}
	})

	self.Connect(self.send, "clicked", func() {
		self.sendTextFromInput()
	})

	self.setupConversation()

}

func (self *ChatWindow) sendTextFromInput() {
	text := self.GetText(self.input)
	if text != "" {
		goline.client.SendText(self.entity.GetId(), text)
		self.input.SetText("")
	}
}

func (self *ChatWindow) setupMessageBox() {
	var err error
	id := self.entity.GetId()
	self.messageBox, err = goline.client.GetMessageBox(id)
	if err != nil {
		goline.LoggerPrintln(err)
		RunAlertMessage(self.window, "Failed to establish connection to server.")
		self.Emit(self.window, "destroy")
	}
}

func (self *ChatWindow) addMessage(message *prot.Message) {
	iMsg := NewIMessage(self, message)
	self.conversation.Add(iMsg.GetIWidget())
	self.conversation.ShowAll()
	self.isAdd = true
}

func (self *ChatWindow) setupConversation() {
	messages, err := goline.client.GetRecentMessages(self.messageBox, 50)
	if err != nil {
		goline.LoggerPrintln(err)
		RunAlertMessage(self.window, "Failed to get recent messages.")
		self.Emit(self.window, "destroy")
	}

	for i := len(messages) - 1; i >= 0; i-- {
		self.addMessage(messages[i])
	}
}

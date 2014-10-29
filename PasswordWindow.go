package main

import (
	"github.com/carylorrk/goline-gotk3/res/glade"
	"github.com/carylorrk/gotk3/gdk"
	"github.com/carylorrk/gotk3/gtk"
)

type PasswordWindow struct {
	*LineWindow
	parent *ConfigWindow
	pwd    *gtk.Entry
	ok     *gtk.Button
}

func NewPasswordWindow(parent *ConfigWindow) *PasswordWindow {
	window := &PasswordWindow{
		LineWindow: NewLineWindow(glade.PasswordWindow),
		parent:     parent}

	window.setupWidgets()
	window.setupDefault()
	return window
}

func (self *PasswordWindow) setupWidgets() {
	self.pwd = self.getIObjectWithType("Password", &gtk.Entry{}).(*gtk.Entry)
	self.ok = self.getIObjectWithType("Ok", &gtk.Button{}).(*gtk.Button)
}

func (self *PasswordWindow) setupDefault() {
	self.Connect(self.window, "key-press-event", func(window *gtk.Window, ev *gdk.Event) {
		keyEvent := &gdk.EventKey{ev}
		keyVal := keyEvent.KeyVal()
		if keyVal == KEY_Return || keyVal == KEY_KP_Enter {
			self.ok.Emit("clicked")
		}
	})

	self.Connect(self.ok, "clicked", func() {
		pwd := self.GetText(self.pwd)
		autologin := self.parent.autologin.GetActive()
		if hashPassword(pwd, goline.Salt) == goline.HashedPassword {
			authToken := goline.AuthToken
			var err error
			if goline.Autologin == false && autologin == true {
				authToken, err = decryptAuthToken(pwd, authToken)
				if err != nil {
					goline.LoggerPrintln(err)
					RunAlertMessage(self.window, "Failed to descrypt authorization token.")
					self.window.Destroy()
				}
			} else if goline.Autologin == true && autologin == false {
				authToken, err = encryptAuthToken(pwd, authToken)
				if err != nil {
					goline.LoggerPrintln(err)
					RunAlertMessage(self.window, "Failed to encrypt authorization token.")
					self.window.Destroy()
				}
			}

			goline.Autologin = autologin
			goline.AuthToken = authToken
			err = goline.SaveSettings()
			if err != nil {
				goline.LoggerPrintln(err)
				RunAlertMessage(self.window, "Failed to update settings.")
			}
			self.parent.window.Destroy()
			self.window.Destroy()
		} else {
			RunAlertMessage(self.window, "Wrong password")
		}
	})
}

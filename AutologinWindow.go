package main

import (
	"github.com/carylorrk/goline-gotk3/res/glade"
	"github.com/carylorrk/gotk3/glib"
	"github.com/carylorrk/gotk3/gtk"
)

type AutologinWindow struct {
	*LineWindow
	exit *gtk.Button
}

func NewAutologinWindow() *AutologinWindow {
	window := &AutologinWindow{LineWindow: NewLineWindow(glade.AutologinWindow)}
	window.setupWidgets()
	return window
}

func (self *AutologinWindow) setupWidgets() {
	self.exit = self.getIObjectWithType("ExitButton", &gtk.Button{}).(*gtk.Button)

	self.Connect(self.window, "destroy", gtk.MainQuit)
	self.Connect(self.exit, "clicked", func() { self.window.Destroy() })
}

func (self *AutologinWindow) Login() {
	errChan := make(chan error)
	go func() {
		err := goline.client.AuthTokenLogin(goline.AuthToken)
		errChan <- err
	}()
	go func() {
		err := <-errChan
		if err != nil {
			goline.LoggerPrintln(err)
			glib.IdleAdd(func() {
				loginWindow := NewLoginWindow()
				loginWindow.window.ShowAll()
				self.window.Hide()
				RunAlertMessage(loginWindow.window, "Failed to login with authorization token.")
			})
		} else {
			glib.IdleAdd(func() {
				NewMainWindow().window.ShowAll()
				self.window.Hide()
			})
		}
	}()
}

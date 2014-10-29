package main

import (
	"github.com/carylorrk/goline/res/glade"
	"github.com/carylorrk/gotk3/glib"
	"github.com/carylorrk/gotk3/gtk"
)

type VerificationWindow struct {
	*LineWindow
	parent *LoginWindow
	code   *gtk.Label
	exit   *gtk.Button
}

func NewVerificationWindow(parent *LoginWindow, pincode string) *VerificationWindow {
	window := &VerificationWindow{
		LineWindow: NewLineWindow(glade.VerificationWindow),
		parent:     parent}

	window.setupWidgets()
	window.setupDefault(pincode)
	window.verifyAndLogin()
	return window
}

func (self *VerificationWindow) setupWidgets() {
	self.code = self.getIObjectWithType("Code", &gtk.Label{}).(*gtk.Label)
	self.exit = self.getIObjectWithType("ExitButton", &gtk.Button{}).(*gtk.Button)

}

func (self *VerificationWindow) setupDefault(pincode string) {
	self.code.SetText(pincode)
	self.window.Connect("destroy", gtk.MainQuit)
	self.Connect(self.exit, "clicked", func() { self.window.Destroy() })
}

func (self *VerificationWindow) verifyAndLogin() {
	errChan := make(chan error)
	authTokenChan := make(chan string)
	go func() {
		authToken, err := goline.client.GetAuthTokenAfterVerify()
		errChan <- err
		authTokenChan <- authToken
	}()

	go func() {
		err := <-errChan
		if err != nil {
			glib.IdleAdd(func() { goline.LoggerPanicln("Failed to get authorization token:", err) })
		}

		authToken := <-authTokenChan
		err = goline.client.AuthTokenLogin(authToken)
		if err != nil {
			glib.IdleAdd(func() { goline.LoggerPanicln("Failed to login with authorization token:", err) })
		}

		glib.IdleAdd(func() {
			mainWindow := NewMainWindow()
			mainWindow.window.ShowAll()
			self.window.Hide()
			if goline.Autologin {
				goline.AuthToken = authToken
			} else {
				goline.AuthToken, err = encryptAuthToken(self.parent.GetText(self.parent.pwd), authToken)
				if err != nil {
					goline.LoggerPrintln(err)
					glib.IdleAdd(func() { RunAlertMessage(mainWindow.window, "Failed to encrypt authorization token.") })
				}
			}
			err = goline.SaveSettings()
			if err != nil {
				goline.LoggerPrintln(err)
				glib.IdleAdd(func() { RunAlertMessage(mainWindow.window, "Failed to save authorization token.") })
			}
		})
	}()
}

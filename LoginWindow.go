package main

import (
	prot "github.com/carylorrk/goline/protocol"
	"github.com/carylorrk/goline/res/glade"
	"github.com/carylorrk/gotk3/gdk"
	"github.com/carylorrk/gotk3/gtk"
)

type LoginWindow struct {
	*LineWindow
	id        *gtk.Entry
	pwd       *gtk.Entry
	autologin *gtk.CheckButton
	login     *gtk.Button
	exit      *gtk.Button
}

func NewLoginWindow() (window *LoginWindow) {
	window = &LoginWindow{LineWindow: NewLineWindow(glade.LoginWindow)}
	window.setupWidgets()
	window.setupDefault()
	return
}

func (self *LoginWindow) setupWidgets() {
	self.id = self.getIObjectWithType("IdEntry", &gtk.Entry{}).(*gtk.Entry)
	self.pwd = self.getIObjectWithType("PasswordEntry", &gtk.Entry{}).(*gtk.Entry)
	self.autologin = self.getIObjectWithType("Autologin", &gtk.CheckButton{}).(*gtk.CheckButton)
	self.login = self.getIObjectWithType("Login", &gtk.Button{}).(*gtk.Button)
	self.exit = self.getIObjectWithType("Exit", &gtk.Button{}).(*gtk.Button)
}

func (self *LoginWindow) setupDefault() {
	self.id.SetText(goline.Id)
	self.autologin.SetActive(goline.Autologin)

	self.window.Connect("destroy", gtk.MainQuit)
	self.Connect(self.window, "key-press-event", func(window *gtk.Window, ev *gdk.Event) {
		keyEvent := &gdk.EventKey{ev}
		keyVal := keyEvent.KeyVal()
		if keyVal == KEY_Return || keyVal == KEY_KP_Enter {
			self.login.Emit("clicked")
		}
	})

	self.Connect(self.login, "clicked", self.handleLogin)
	self.Connect(self.exit, "clicked", func() { self.window.Destroy() })
}

func (self *LoginWindow) handleLogin() {
	self.login.SetSensitive(false)
	defer self.login.SetSensitive(true)
	id := self.GetText(self.id)
	pwd := self.GetText(self.pwd)

	if !goline.Autologin && id == goline.Id && goline.AuthToken != "" {
		self.loginWithDecryptedAuthToken(pwd)
	} else {
		self.loginWithVerification(id, pwd)
	}
}

func (self *LoginWindow) loginWithDecryptedAuthToken(pwd string) {
	if hashPassword(pwd, goline.Salt) != goline.HashedPassword {
		goline.LoggerPrintln("Failed to login with wrong password.")
		RunAlertMessage(self.window, "Failed to login with wrong password.")
		return
	}

	authToken, err := decryptAuthToken(pwd, goline.AuthToken)
	if err != nil {
		goline.LoggerPrintln(err)
		RunAlertMessage(self.window, "Failed to decrypt authorization token.")
		goline.AuthToken = ""
		err = goline.SaveSettings()
		if err != nil {
			goline.LoggerPrintln(err)
			RunAlertMessage(self.window, "Failed to clean authorization token.")
		}
		return
	}
	err = goline.client.AuthTokenLogin(authToken)
	if err != nil {
		goline.LoggerPrintln(err)
		RunAlertMessage(self.window, "Failed to login with autorization token.")
		goline.AuthToken = ""
		err = goline.SaveSettings()
		if err != nil {
			goline.LoggerPrintln(err)
			RunAlertMessage(self.window, "Failed to clean authorization token.")
		}
		return
	}
	if self.autologin.GetActive() {
		goline.AuthToken = authToken
	}
	self.updateSettings()
	NewMainWindow().window.ShowAll()
	self.window.Hide()

}

func (self *LoginWindow) updateSettings() {
	goline.Autologin = self.autologin.GetActive()
	goline.Id = self.GetText(self.id)
	goline.Salt = generateRandomString()
	goline.HashedPassword = hashPassword(self.GetText(self.pwd), goline.Salt)
	err := goline.SaveSettings()
	if err != nil {
		goline.LoggerPrintln(err)
		RunAlertMessage(self.window, "Failed to update settings.")
	}
}

func (self *LoginWindow) loginWithVerification(id, pwd string) {
	goline.AuthToken = ""
	self.updateSettings()
	pincode, err := self.getPincode(id, pwd)
	if err != nil {
		goline.LoggerPrintln(err)
		var reason string
		switch v := err.(type) {
		case *prot.TalkException:
			reason = v.GetReason()
		default:
			reason = "Ooooops, something went wrong!"
		}
		RunAlertMessage(self.window, reason)
		return
	}
	goline.LoggerPrintln("pincode for", id, ":", pincode)
	self.verify(pincode)
}

func (self *LoginWindow) getPincode(id, passwd string) (string, error) {
	pincode, err := goline.client.GetPincode(id, passwd)
	if err != nil {
		goline.LoggerPrintln(err)
		return "", err
	}
	return pincode, nil
}

func (self *LoginWindow) verify(pincode string) {
	NewVerificationWindow(self, pincode).window.ShowAll()
	self.window.Hide()
}

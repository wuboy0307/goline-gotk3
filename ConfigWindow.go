package main

import (
	"github.com/carylorrk/goline/res/glade"
	"github.com/carylorrk/gotk3/gtk"
)

type ConfigWindow struct {
	*LineWindow
	dataDir   *gtk.Entry
	tempDir   *gtk.Entry
	autologin *gtk.CheckButton
	save      *gtk.Button
	cancel    *gtk.Button
}

func NewConfigWindow() *ConfigWindow {
	window := &ConfigWindow{LineWindow: NewLineWindow(glade.ConfigWindow)}
	window.setupWidgets()
	window.setupDefault()
	return window
}

func (self *ConfigWindow) setupWidgets() {
	self.dataDir = self.getIObjectWithType("DataDirEntry", &gtk.Entry{}).(*gtk.Entry)
	self.tempDir = self.getIObjectWithType("TempDirEntry", &gtk.Entry{}).(*gtk.Entry)
	self.autologin = self.getIObjectWithType("Autologin", &gtk.CheckButton{}).(*gtk.CheckButton)
	self.save = self.getIObjectWithType("Save", &gtk.Button{}).(*gtk.Button)
	self.cancel = self.getIObjectWithType("Cancel", &gtk.Button{}).(*gtk.Button)

}

func (self *ConfigWindow) setupDefault() {
	self.dataDir.SetText(goline.DataDirPath)
	self.tempDir.SetText(goline.TempDirPath)
	self.autologin.SetActive(goline.Autologin)

	self.Connect(self.save, "clicked", func() {
		NewPasswordWindow(self).window.ShowAll()
	})
	self.Connect(self.cancel, "clicked", self.window.Destroy)

}

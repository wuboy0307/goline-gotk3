package main

import (
	"os/exec"

	"github.com/conformal/gotk3/glib"
	"github.com/conformal/gotk3/gtk"
)

type PopupMenu struct {
	parent   *ChatWindow
	menu     *gtk.Menu
	view     *gtk.MenuItem
	download *gtk.MenuItem
}

func NewPopupMenu(parent *ChatWindow, msgId, url, ext string) *PopupMenu {
	menu := &PopupMenu{parent: parent}
	menu.setupWidgets()
	menu.setupDefault(msgId, url, ext)
	return menu
}

func (self *PopupMenu) setupWidgets() {
	var err error
	self.menu, err = gtk.MenuNew()
	if err != nil {
		goline.LoggerPanicln(err)
	}

	self.view, err = gtk.MenuItemNew()
	if err != nil {
		goline.LoggerPanicln(err)
	}
	self.view.SetLabel("View")

	self.download, err = gtk.MenuItemNew()
	if err != nil {
		goline.LoggerPanicln(err)
	}
	self.download.SetLabel("Download")

	self.menu.Append(self.view)
	self.menu.Append(self.download)

	self.menu.ShowAll()
}

func (self *PopupMenu) setupDefault(msgId, url, ext string) {
	_, err := self.view.Connect("activate", func() {
		self.viewFile(msgId, url, ext)
	})
	if err != nil {
		goline.LoggerPanicln(err)
	}

	_, err = self.download.Connect("activate", func() {
		self.downloadFile(msgId, url, ext)
	})
	if err != nil {
		goline.LoggerPanicln(err)
	}
}

func (self *PopupMenu) downloadFile(msgId, url, ext string) {
	go func() {
		glib.IdleAdd(func() {
			NewFileChooserWindow(self.parent, msgId, url, ext).window.ShowAll()
		})
	}()
}

func (self *PopupMenu) viewFile(msgId, url, ext string) {
	go func() {
		downloadingWindow := NewDownloadingWindow(msgId)
		glib.IdleAdd(downloadingWindow.window.ShowAll)
		defer glib.IdleAdd(downloadingWindow.window.Destroy)
		content, err := downloadContentToTemp(msgId, url, ext)
		if err != nil {
			goline.LoggerPrintln(err)
			glib.IdleAdd(func() { RunAlertMessage(self.parent.window, "Failed to download file: %s", err) })
			return
		}

		cmd := exec.Command("xdg-open", content)
		err = cmd.Start()
		if err != nil {
			goline.LoggerPrintln(err)
			glib.IdleAdd(func() { RunAlertMessage(self.parent.window, "Failed to open file.") })
			return
		}
	}()
}

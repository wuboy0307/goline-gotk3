package main

import (
	"github.com/carylorrk/goline-gotk3/res/glade"
	"github.com/carylorrk/gotk3/gtk"
)

type DownloadableMessage struct {
	*Message
	preview    gtk.IWidget
	menu       *gtk.MenuButton
	previewBox *gtk.EventBox
}

func NewDownloadableMessage(parent *ChatWindow, msgType MessageType, from, msgId, url, ext string, preview gtk.IWidget) *DownloadableMessage {
	msg := &DownloadableMessage{
		Message: NewMessage(parent, msgType, from, glade.DownloadableContent),
		preview: preview}

	msg.setupWidgets()
	msg.setupDefault(msgId, url, ext)
	return msg
}

func (self *DownloadableMessage) setupWidgets() {
	self.menu = self.getIObjectWithType("MenuButton", &gtk.MenuButton{}).(*gtk.MenuButton)
	self.previewBox = self.getIObjectWithType("PreviewBox", &gtk.EventBox{}).(*gtk.EventBox)

	self.previewBox.Add(self.preview)
}

func (self *DownloadableMessage) setupDefault(msgId, url, ext string) {
	self.menu.SetPopup(NewPopupMenu(self.parent, msgId, url, ext).menu)
}

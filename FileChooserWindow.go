package main

import (
	"github.com/carylorrk/goline-gotk3/res/glade"
	"github.com/conformal/gotk3/gtk"
	"github.com/mattn/go-gtk/glib"
	"os"
	"path"
)

type FileChooserWindow struct {
	*LineWindow
	parent      *ChatWindow
	fileChooser *gtk.FileChooserWidget
	save        *gtk.Button
	cancel      *gtk.Button
}

func NewFileChooserWindow(parent *ChatWindow, msgId, url, ext string) *FileChooserWindow {
	window := &FileChooserWindow{
		parent:     parent,
		LineWindow: NewLineWindow(glade.FileChooserWindow)}
	window.setupWidgets()
	window.setupDefult(msgId, url, ext)
	return window
}

func (self *FileChooserWindow) setupWidgets() {
	self.fileChooser = self.getIObjectWithType("FileChooser", &gtk.FileChooserWidget{}).(*gtk.FileChooserWidget)
	self.save = self.getIObjectWithType("Save", &gtk.Button{}).(*gtk.Button)
	self.cancel = self.getIObjectWithType("Cancel", &gtk.Button{}).(*gtk.Button)
}

func (self *FileChooserWindow) setupDefult(msgId, url, ext string) {
	self.Connect(self.save, "clicked", func() {
		dirname := self.fileChooser.GetFilename()
		if dirname != "" {
			go func() {
				glib.IdleAdd(self.window.Destroy)
				downloadingWindow := NewDownloadingWindow(msgId)
				glib.IdleAdd(downloadingWindow.window.ShowAll)
				defer glib.IdleAdd(downloadingWindow.window.Destroy)
				content, err := downloadContentToTemp(msgId, url, ext)
				if err != nil {
					goline.LoggerPrintln(err)
					glib.IdleAdd(func() { RunAlertMessage(self.parent.window, "Failed to download file: %s", err) })
					return
				}
				err = os.Rename(content, path.Join(dirname, path.Base(content)))
				if err != nil {
					goline.LoggerPrintln(err)
					glib.IdleAdd(func() { RunAlertMessage(self.parent.window, "Failed to move file: %s", err) })
					return
				}

			}()
		}
	})
	self.Connect(self.cancel, "clicked", func() {
		glib.IdleAdd(self.window.Destroy)
	})
}

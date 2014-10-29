package main

import (
	"github.com/carylorrk/goline/res/glade"
	"github.com/carylorrk/gotk3/gtk"
)

type DownloadingWindow struct {
	*LineWindow
	filename *gtk.Label
}

func NewDownloadingWindow(filename string) *DownloadingWindow {
	window := &DownloadingWindow{LineWindow: NewLineWindow(glade.DownloadingWindow)}
	window.filename = window.getIObjectWithType("Filename", &gtk.Label{}).(*gtk.Label)
	window.filename.SetText(filename)
	return window
}

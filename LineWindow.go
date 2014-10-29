package main

import (
	"github.com/carylorrk/goline-gotk3/res/image"
	"github.com/carylorrk/gotk3/gdk"
	"github.com/carylorrk/gotk3/glib"
	"github.com/carylorrk/gotk3/gtk"
)

type LineWindow struct {
	builder *gtk.Builder
	window  *gtk.Window
}

func NewLineWindow(glade string) (window *LineWindow) {
	window = &LineWindow{}
	window.setupBuilder(glade)
	window.setupWidgets()
	window.setupIcon(image.Icon)

	return window
}

func (self *LineWindow) setupBuilder(glade string) {
	var err error
	self.builder, err = gtk.BuilderNew()
	if err != nil {
		goline.LoggerPanicln(err)
	}

	err = self.builder.AddFromString(glade)
	if err != nil {
		goline.LoggerPanicln(err)
	}
}

func (self *LineWindow) setupWidgets() {
	self.window = self.getIObjectWithType("Window", &gtk.Window{}).(*gtk.Window)
}

func (self *LineWindow) setupIcon(icon []byte) {
	var err error
	iconPixbuf, err := gdk.PixbufNewFromData(
		icon, gdk.COLORSPACE_RGB, false,
		8, 200, 200, 200*3)
	if err != nil {
		goline.LoggerPanicln(err)
	}

	self.window.SetIcon(iconPixbuf)
	if err != nil {
		goline.LoggerPanicln(err)
	}
}

func (self *LineWindow) getIObjectWithType(name string, proto interface{}) glib.IObject {
	return getIObjectWithType(self.builder, name, proto)
}

type Connectable interface {
	Connect(string, interface{}, ...interface{}) (glib.SignalHandle, error)
}

func (self *LineWindow) Connect(obj Connectable, detailedSignal string, f interface{}, userData ...interface{}) glib.SignalHandle {
	handler, err := obj.Connect(detailedSignal, f, userData...)
	if err != nil {
		goline.LoggerPanicln(err)
	}
	return handler
}

type Emitable interface {
	Emit(s string, args ...interface{}) (interface{}, error)
}

func (self *LineWindow) Emit(obj Emitable, s string, args ...interface{}) interface{} {
	res, err := obj.Emit(s, args...)
	if err != nil {
		goline.LoggerPanicln(err)
	}
	return res
}

type TextGetable interface {
	GetText() (string, error)
}

func (self *LineWindow) GetText(obj TextGetable) string {
	text, err := obj.GetText()
	if err != nil {
		goline.LoggerPanicln(err)
	}
	return text
}

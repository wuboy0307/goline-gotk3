package main

import (
	"github.com/carylorrk/goline/res/glade"
	"github.com/carylorrk/gotk3/glib"
	"github.com/carylorrk/gotk3/gtk"
)

type MessageType int

const (
	MESSAGE_NORMAL MessageType = 0
	MESSAGE_YOU    MessageType = 1
)

type Message struct {
	msgBuilder *gtk.Builder
	builder    *gtk.Builder

	message   *gtk.Box
	alignment *gtk.Alignment
	name      *gtk.Label
	content   *gtk.Box

	parent *ChatWindow
}

func NewMessage(parent *ChatWindow, msgType MessageType, from, gladeStr string) *Message {
	msg := &Message{parent: parent}
	msg.setupBuilder(msgType, gladeStr)
	msg.setupWidges()
	msg.setupDefault(msgType, from)
	return msg
}

func (self *Message) setupDefault(msgType MessageType, from string) {
	if msgType == MESSAGE_NORMAL {
		name := getNameById(from)
		self.name.SetMarkup(markupById(from, name))
	}
}

func (self *Message) setupBuilder(msgType MessageType, gladeStr string) {
	var err error
	self.msgBuilder, err = gtk.BuilderNew()
	if err != nil {
		goline.LoggerPanicln(err)
	}

	switch msgType {
	case MESSAGE_NORMAL:
		err = self.msgBuilder.AddFromString(glade.Message)
	case MESSAGE_YOU:
		err = self.msgBuilder.AddFromString(glade.YouMessage)
	}

	if err != nil {
		goline.LoggerPanicln(err)
	}

	self.builder, err = gtk.BuilderNew()
	if err != nil {
		goline.LoggerPanicln(err)
	}

	err = self.builder.AddFromString(gladeStr)
	if err != nil {
		goline.LoggerPanicln(err)
	}
}

func (self *Message) setupWidges() {
	self.message = getIObjectWithType(self.msgBuilder, "Message", &gtk.Box{}).(*gtk.Box)
	self.alignment = getIObjectWithType(self.msgBuilder, "Alignment", &gtk.Alignment{}).(*gtk.Alignment)
	self.name = getIObjectWithType(self.msgBuilder, "Name", &gtk.Label{}).(*gtk.Label)
	self.content = self.getIObjectWithType("Content", &gtk.Box{}).(*gtk.Box)

	self.alignment.Add(self.content)
}

func (self *Message) getIObjectWithType(name string, proto interface{}) glib.IObject {
	return getIObjectWithType(self.builder, name, proto)
}

func (self *Message) GetIWidget() gtk.IWidget {
	return self.message
}

func (self *Message) Connect(obj Connectable, detailedSignal string, f interface{}, userData ...interface{}) glib.SignalHandle {
	handler, err := obj.Connect(detailedSignal, f, userData...)
	if err != nil {
		goline.LoggerPanicln(err)
	}
	return handler
}

package main

import (
	"github.com/carylorrk/goline-gotk3/res/glade"
	"github.com/conformal/gotk3/gtk"
)

type TextMessage struct {
	*Message
	text *gtk.Label
}

func NewTextMessage(parent *ChatWindow, msgType MessageType, from, text string) *TextMessage {
	msg := &TextMessage{Message: NewMessage(parent, msgType, from, glade.TextContent)}
	msg.setupWidgets()
	msg.setupDefault(msgType, from, text)
	return msg
}

func (self *TextMessage) setupWidgets() {
	self.text = self.getIObjectWithType("Text", &gtk.Label{}).(*gtk.Label)
}

func (self *TextMessage) setupDefault(msgType MessageType, from, text string) {
	if msgType == MESSAGE_NORMAL {
		self.text.SetMarkup(markupById(from, text))
	} else {
		self.text.SetText(text)
	}
}

package main

import (
	"github.com/carylorrk/goline/res/glade"
	"github.com/carylorrk/gotk3/gtk"
)

type ImageMessage struct {
	*Message
	image *gtk.Image
}

func NewImageMessage(parent *ChatWindow, msgType MessageType, from, filename string) *ImageMessage {
	msg := &ImageMessage{Message: NewMessage(parent, msgType, from, glade.ImageContent)}

	msg.image = msg.getIObjectWithType("Image", &gtk.Image{}).(*gtk.Image)
	msg.image.SetFromFile(filename)
	return msg
}
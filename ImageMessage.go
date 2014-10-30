package main

import (
	"github.com/carylorrk/goline-gotk3/res/glade"
	"github.com/conformal/gotk3/gtk"
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

package main

import (
	"path"

	"github.com/carylorrk/goline/api"
	prot "github.com/carylorrk/goline/protocol"
	"github.com/carylorrk/gotk3/gtk"
)

type IMessage interface {
	GetIWidget() gtk.IWidget
}

func checkMessageType(from string) MessageType {
	if from == goline.client.Profile.GetMid() {
		return MESSAGE_YOU
	} else {
		return MESSAGE_NORMAL
	}

}

func NewIMessage(parent *ChatWindow, msg *prot.Message) IMessage {
	contentType := msg.GetContentType()

	//TODO: Support content type
	switch contentType {
	case prot.ContentType_NONE:
		return handleTextMessage(parent, msg, msg.GetText())
	case prot.ContentType_IMAGE:
		return handleImagePreviewMessage(parent, msg)
	case prot.ContentType_VIDEO:
		return handleImagePreviewMessage(parent, msg)
	case prot.ContentType_STICKER:
		return handleStickerMessage(parent, msg)
	case prot.ContentType_AUDIO:
		return handleTextPreviewMessage(parent, msg, "Content Type: Audio", ".mp4")
	case prot.ContentType_FILE:
		return handleTextPreviewMessage(parent, msg, "Content Type: File", "")
	default:
		return handleTextMessage(parent, msg, contentType.String())
	}
}

func handleTextMessage(parent *ChatWindow, msg *prot.Message, text string) IMessage {
	from := msg.GetFrom()

	msgType := checkMessageType(from)
	return NewTextMessage(parent, msgType, from, text)
}

func handleStickerMessage(parent *ChatWindow, msg *prot.Message) IMessage {
	from := msg.GetFrom()
	msgType := checkMessageType(from)

	meta := msg.ContentMetadata
	stkid := meta["STKID"]
	stkpkgid := meta["STKPKGID"]
	stkver := meta["STKVER"]
	url := api.LINE_STICKER_URL + stkver + "/" + stkpkgid + "/PC/stickers/" + stkid + ".png"
	filename := path.Join(goline.TempDirPath, "sticker", stkid+".png")
	if checkFileNotExist(filename) {
		err := downloadFile(url, filename)
		if err != nil {
			goline.LoggerPrintln(err)
			return NewTextMessage(parent, msgType, from, "Failed to download sticker.")

		}
	}
	return NewImageMessage(parent, msgType, from, filename)
}

func handleImagePreviewMessage(parent *ChatWindow, msg *prot.Message) IMessage {
	from := msg.GetFrom()
	msgType := checkMessageType(from)

	msgId := msg.GetId()
	preview := path.Join(goline.TempDirPath, "preview", msgId)
	meta := msg.GetContentMetadata()

	var url string
	if checkFileNotExist(preview) {
		if meta["PUBLIC"] == "TRUE" {
			url = meta["PREVIEW_URL"]
		} else {
			url = api.LINE_OBJECT_STORAGE_URL + msgId + "/preview"
		}
		err := downloadFile(url, preview)
		if err != nil {
			goline.LoggerPrintln(err)
			return NewTextMessage(parent, msgType, from, "Failed to download preview.")
		}
	}

	image, err := gtk.ImageNew()
	if err != nil {
		goline.LoggerPanicln(err)
	}
	image.SetFromFile(preview)

	if meta["PUBLIC"] == "TRUE" {
		url = meta["DOWNLOAD_URL"]
	} else {
		url = api.LINE_OBJECT_STORAGE_URL + msgId
	}

	return NewDownloadableMessage(parent, msgType, from, msgId, url, "", image)
}

func handleTextPreviewMessage(parent *ChatWindow, msg *prot.Message, text, ext string) IMessage {
	from := msg.GetFrom()
	msgType := checkMessageType(from)

	label, err := gtk.LabelNew(text)
	if err != nil {
		goline.LoggerPanicln(err)
	}
	msgId := msg.GetId()
	meta := msg.GetContentMetadata()

	var url string
	if meta["PUBLIC"] == "TRUE" {
		url = meta["DOWNLOAD_URL"]
	} else {
		url = api.LINE_OBJECT_STORAGE_URL + msgId
	}

	return NewDownloadableMessage(parent, msgType, from, msgId, url, ext, label)

}

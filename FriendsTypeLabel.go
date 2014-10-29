package main

import (
	"github.com/carylorrk/goline/res/glade"
	"github.com/carylorrk/gotk3/gtk"
)

type FriendsTypeLabel struct {
	builder *gtk.Builder
	label   *gtk.Label
}

func NewFriendsTypeLabel(text string) *FriendsTypeLabel {
	builder, err := gtk.BuilderNew()
	if err != nil {
		goline.LoggerPanicln(err)
	}
	err = builder.AddFromString(glade.FriendsTypeLabel)
	if err != nil {
		goline.LoggerPanicln(err)
	}

	label := getIObjectWithType(builder, "FriendsType", &gtk.Label{}).(*gtk.Label)
	label.SetText(text)
	return &FriendsTypeLabel{builder: builder, label: label}
}

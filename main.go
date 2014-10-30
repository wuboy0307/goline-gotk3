package main

import (
	"fmt"
	"github.com/carylorrk/goline-gotk3/res/css"
	"github.com/conformal/gotk3/gtk"
	"os"
)

var goline *Goline

func init() {
	var err error
	goline, err = NewGoline()
	if err != nil {
		panic(fmt.Sprintf("NewGoline Error: %s", err))
	}
}

func main() {
	goline.LoggerPrintln("Start Goline.")
	gtk.Init(&os.Args)
	css.SetupCss()
	if goline.Autologin && goline.AuthToken != "" {
		autologinWindow := NewAutologinWindow()
		autologinWindow.window.ShowAll()
		autologinWindow.Login()
	} else {
		loginWindow := NewLoginWindow()
		loginWindow.window.ShowAll()
	}
	gtk.Main()
}

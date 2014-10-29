package main

import (
	"encoding/json"
	"fmt"
	"github.com/carylorrk/goline-gotk3/api"
	"io"
	"log"
	"os"
	"os/user"
	"path"
	"runtime"
)

type Goline struct {
	Id             string          `json:"Id"`
	HashedPassword string          `json:"HashedPassword"`
	Salt           string          `json:Salt`
	AuthToken      string          `json:"AuthToken"`
	Autologin      bool            `json:"Autologin"`
	DataDirPath    string          `json:"-"`
	TempDirPath    string          `json:"-"`
	client         *api.LineClient `json:"-"`
	logger         *log.Logger     `json:"-"`
}

func NewGoline() (goline *Goline, err error) {
	goline = &Goline{}
	err = goline.setupDirPath()
	if err != nil {
		return
	}

	err = goline.setupLogger()
	if err != nil {
		return
	}

	err = goline.setupSettings()
	if err != nil {
		goline.LoggerPrintln(err)
		return
	}
	goline.client, err = api.NewLineClient()
	if err != nil {
		goline.LoggerPrintln(err)
		return
	}
	return
}

func (self *Goline) setupDirPath() (err error) {
	var usr *user.User
	usr, err = user.Current()
	if err != nil {
		return
	}
	self.DataDirPath = path.Join(usr.HomeDir, ".goline")
	err = os.MkdirAll(self.DataDirPath, os.FileMode(0700))
	if err != nil {
		return
	}

	golineTempDir := path.Join(os.TempDir(), "goline")
	if checkFileNotExist(golineTempDir) {
		err = os.MkdirAll(golineTempDir, os.ModePerm)
		if err != nil {
			return
		}

		err = os.Chmod(golineTempDir, os.ModePerm)
		if err != nil {
			return
		}
	}

	self.TempDirPath = path.Join(golineTempDir, usr.Username)
	err = os.MkdirAll(self.TempDirPath, os.FileMode(0700))
	if err != nil {
		return
	}

	previewPath := path.Join(self.TempDirPath, "preview")
	err = os.MkdirAll(previewPath, os.FileMode(0700))
	if err != nil {
		return
	}

	contentPath := path.Join(self.TempDirPath, "content")
	err = os.MkdirAll(contentPath, os.FileMode(0700))
	if err != nil {
		return
	}

	stickerPath := path.Join(self.TempDirPath, "sticker")
	err = os.MkdirAll(stickerPath, os.FileMode(0700))
	if err != nil {
		return
	}

	thumbnailPath := path.Join(self.TempDirPath, "thumbnail")
	err = os.MkdirAll(thumbnailPath, os.FileMode(0700))
	if err != nil {
		return
	}

	imagePath := path.Join(self.TempDirPath, "image")
	err = os.MkdirAll(imagePath, os.FileMode(0700))
	return
}

func (self *Goline) loadConfigFile() (*os.File, error) {
	configFilePath := path.Join(self.DataDirPath, "settings2.json")
	configFile, err := os.OpenFile(configFilePath, os.O_RDWR|os.O_CREATE, os.FileMode(0600))
	if err != nil {
		return nil, err
	}
	return configFile, err
}

func (self *Goline) SaveSettings() error {
	configFile, err := self.loadConfigFile()
	if err != nil {
		return err
	}
	defer configFile.Close()
	jsonEncoder := json.NewEncoder(configFile)
	err = jsonEncoder.Encode(self)
	return err
}

func (self *Goline) setupSettings() error {
	configFile, err := self.loadConfigFile()
	if err != nil {
		return err
	}
	defer configFile.Close()
	jsonDecoder := json.NewDecoder(configFile)
	err = jsonDecoder.Decode(self)
	if err != nil {
		err = configFile.Truncate(0)
		if err != nil {
			return err
		}
		self.logger.Println("Create new setting file.")
		jsonEncoder := json.NewEncoder(configFile)
		err = jsonEncoder.Encode(self)
		if err != nil {
			return err
		}
	}
	return err
}

func (self *Goline) setupLogger() error {
	logFilePath := path.Join(self.TempDirPath, "log")
	logFile, err := os.OpenFile(logFilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		return err
	}
	writer := io.MultiWriter(logFile, os.Stderr)
	self.logger = log.New(writer, "Goline: ", log.LstdFlags)
	return nil
}

func (self *Goline) LoggerPrintln(v ...interface{}) {
	pc, file, line, ok := runtime.Caller(1)
	file = path.Base(file)
	if ok {
		self.logger.Println(pc, file, line, v)
	} else {
		self.logger.Println(v)
	}
}

func (self *Goline) LoggerPanicln(v ...interface{}) {
	pc, file, line, ok := runtime.Caller(1)
	file = path.Base(file)
	if ok {
		msg := fmt.Sprintln(pc, file, line, v)
		RunErrorMessage(msg)
		self.logger.Panicln(pc, file, line, v)
	} else {
		RunErrorMessage(fmt.Sprintln(v))
		self.logger.Panicln(v)
	}
}

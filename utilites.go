package main

import (
	"crypto/aes"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	"html"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path"
	"reflect"
	"time"

	"github.com/carylorrk/gotk3/glib"
	"github.com/carylorrk/gotk3/gtk"
)

func checkFileNotExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return true
	}
	return false
}

func getResponse(url string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		goline.LoggerPrintln(err)
		return nil, err
	}
	req.Header = *goline.client.GetHeader()
	res, err := client.Do(req)
	if err != nil {
		goline.LoggerPrintln(err)
		return nil, err
	}
	return res, nil
}

func copyBodyToFile(res *http.Response, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		goline.LoggerPrintln(err)
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, res.Body)
	if err != nil {
		goline.LoggerPrintln(err)
	}
	return err
}

func downloadFile(url, filename string) error {
	res, err := getResponse(url)
	if err != nil {
		goline.LoggerPrintln(err)
		return err
	}
	defer res.Body.Close()
	err = copyBodyToFile(res, filename)
	if err != nil {
		goline.LoggerPrintln(err)
	}
	return err
}

func downloadContentToTemp(msgId, url, ext string) (string, error) {
	res, err := getResponse(url)
	defer res.Body.Close()
	if err != nil {
		goline.LoggerPrintln(err)
		return "", err
	}
	if ext == "" {
		ext = mimeToExt[res.Header.Get("Content-Type")]
	}
	content := path.Join(goline.TempDirPath, "content", msgId) + ext
	goline.LoggerPrintln("Save file to", content)
	err = copyBodyToFile(res, content)
	if err != nil {
		goline.LoggerPrintln(err)
		return "", err
	}
	return content, nil
}

func getNameById(id string) string {
	entity, err := goline.client.GetLineEntityById(id)
	if err != nil || entity == nil {
		return "Unknown"
	} else {
		return entity.GetName()
	}
}

func generateRandomColorFromId(id string) string {
	num := []byte(id)
	var (
		r uint8 = 20
		g uint8 = 20
		b uint8 = 20
	)
	if uint8(num[0])%2 == 1 {
		r += uint8(num[10]) % 100
	}
	if uint8(num[1]%2) == 1 {
		g += uint8(num[11]) % 100
	}
	if uint8(num[2]%2) == 1 {
		b += uint8(num[12]) % 100
	}

	return fmt.Sprintf("#%x%x%x", r, g, b)
}

func markupById(id, text string) string {
	color := generateRandomColorFromId(id)
	format := "<span foreground=\"%s\">%s</span>"
	return fmt.Sprintf(format,
		html.EscapeString(color),
		html.EscapeString(text))
}

func getIObjectWithType(builder *gtk.Builder, name string, proto interface{}) glib.IObject {
	iobj, err := builder.GetObject(name)
	if err != nil {
		goline.LoggerPanicln(err)
	}
	protoType := reflect.TypeOf(proto)
	iobjType := reflect.TypeOf(iobj)
	if iobjType != protoType {
		goline.LoggerPanicln(NewTypeAssertionError(name, protoType.Name(), iobjType.Name()))
	}

	return iobj
}

func generateRandomString() string {
	rand.Seed(time.Now().UnixNano())
	n := rand.Int()%69 + 32
	s := make([]byte, n)
	for i := 0; i < n; i++ {
		x := rand.Int()%95 + 32
		s[i] = byte(x)
	}
	return string(s)

}

func hashPassword(pwd, salt string) string {
	sum := sha1.Sum([]byte(pwd + salt))
	return base64.StdEncoding.EncodeToString(sum[:])
}

func encryptAuthToken(pwd string, authToken string) (string, error) {
	sum := md5.Sum([]byte(pwd))
	key := sum[0:16]
	block, err := aes.NewCipher(key)
	if err != nil {
		goline.LoggerPrintln(err)
		return "", err
	}
	authTokenBytes := []byte(authToken)
	block.Encrypt(authTokenBytes, authTokenBytes)
	authToken = base64.StdEncoding.EncodeToString(authTokenBytes)
	return authToken, nil
}

func decryptAuthToken(pwd string, authTokenStr string) (string, error) {
	authToken, err := base64.StdEncoding.DecodeString(authTokenStr)
	if err != nil {
		goline.LoggerPrintln(err)
		return "", err
	}
	sum := md5.Sum([]byte(pwd))
	key := sum[0:16]
	block, err := aes.NewCipher(key)
	if err != nil {
		goline.LoggerPanicln(err)
		return "", err
	}
	block.Decrypt(authToken, authToken)
	return string(authToken), nil
}

func NewTypeAssertionError(name, want string, get string) error {
	return errors.New(fmt.Sprintf("%s is not a %s but a %s.", name, want, get))
}

func RunAlertMessage(parent *gtk.Window, format string, args ...interface{}) {
	dialog := gtk.MessageDialogNew(parent, gtk.DIALOG_MODAL, gtk.MESSAGE_WARNING, gtk.BUTTONS_OK, fmt.Sprint("WARNING: ", format), args...)
	dialog.Run()
	dialog.Destroy()
}

func RunErrorMessage(format string, args ...interface{}) error {
	dummy, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		return err
	}
	dialog := gtk.MessageDialogNew(dummy, gtk.DIALOG_MODAL, gtk.MESSAGE_ERROR, gtk.BUTTONS_OK, fmt.Sprint("ERROR: ", format), args...)
	dialog.Run()
	dialog.Destroy()
	return nil
}

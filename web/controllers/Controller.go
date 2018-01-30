package controllers

import (
	"html/template"
	"net/http"
	"github.com/gorilla/sessions"
	"fmt"
	"math/rand"
	"reflect"
	"strings"
	"os"
)

var (
	store = sessions.NewCookieStore([]byte(randSeq(10)))
	letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

type flash struct {
	Message string
	Level string
}

func renderView(w http.ResponseWriter, r *http.Request, baseview string, view string, data interface{}) {
	type d struct {
		Data interface{}
		Flash flash
		Version string
	}
	t := template.Must(template.ParseFiles("./web/views/layouts/"+baseview, "./web/views/"+view))
	t.ExecuteTemplate(w, "layout", d{
		Data:data,
		Flash:getFlash(w, r),
		Version: os.Getenv("VERSION"),
	})
}

func setFlash(message string, level string, w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "flash-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	session.AddFlash(message, level)
	session.Save(r, w)
}

func getFlash(w http.ResponseWriter, r *http.Request) (message flash) {
	session, err := store.Get(r, "flash-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fm := session.Values
	if fm == nil {
		return
	}
	session.Save(r, w)

	session.Options.MaxAge = -1
	session.Save(r, w)
	//return fmt.Sprintf("%v", fm[0])
	keys := reflect.ValueOf(session.Values).MapKeys()
	if len(keys) < 1 {
		return
	}
	key := keys[0].Interface()
	value := fmt.Sprintf("%s", session.Values[key])
	value = strings.Replace(value, "[", "", -1)
	value = strings.Replace(value, "]", "", -1)

	return flash{
		Message:value,
		Level:fmt.Sprintf("%v", key),
	}
}

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

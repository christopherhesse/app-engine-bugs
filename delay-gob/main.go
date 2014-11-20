package hello

import (
	"bytes"
	"encoding/gob"
	"net/http"

	"appengine"
	"appengine/delay"
)

func init() {
	gob.RegisterName("hello.Object", Object{})
	http.HandleFunc("/", handler)
}

type Object struct {
	String string
}

func handler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	buf := &bytes.Buffer{}
	args := []interface{}{Object{}}
	gob.NewEncoder(buf).Encode(args)
	c.Infof("payload=%s", buf.Bytes())
}

var laterFunc = delay.Func("key", func(c appengine.Context, obj Object) {
	c.Infof("later")
})

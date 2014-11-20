package hello

import (
	"bytes"
	"encoding/gob"
	"net/http"

	"appengine"
	"appengine/delay"
)

func init() {
	http.HandleFunc("/", handler)
}

type Object struct {
	String string
}

func handler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	t, err := laterFunc.Task(Object{})
	c.Infof("err=%+v", err)
	c.Infof("payload=%s", t.Payload)

	buf := &bytes.Buffer{}
	args := []interface{}{Object{}}
	err = gob.NewEncoder(buf).Encode(args)
	c.Infof("err=%+v", err)
	c.Infof("payload=%s", buf.Bytes())
}

var laterFunc = delay.Func("key", func(c appengine.Context, obj Object) {
	c.Infof("later")
})

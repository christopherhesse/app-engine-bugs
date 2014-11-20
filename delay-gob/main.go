package hello

import (
	"errors"
	"net/http"

	"submodule"

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
	c.Infof("%T %T", Object{}, submodule.Object{})

	t, err := laterFunc.Task(Object{}, submodule.Object{})
	c.Infof("err=%+v", err)
	c.Infof("payload=%s", t.Payload)
}

var laterFunc = delay.Func("key", func(c appengine.Context, obj1 Object, obj2 submodule.Object) error {
	return errors.New("retry")
})

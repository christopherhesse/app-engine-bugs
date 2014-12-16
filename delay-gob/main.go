package hello

import (
	"bytes"
	"encoding/gob"
	"net/http"
	"reflect"
	"regexp"

	"sub/submodule"

	"appengine"
	"appengine/delay"
)

func init() {
	http.HandleFunc("/", handler)
}

type Object struct {
	String string
}

// based on Register() from https://code.google.com/p/go/source/browse/src/encoding/gob/type.go
func Name(value interface{}) string {
	rt := reflect.TypeOf(value)
	name := rt.String()
	star := ""
	if rt.Name() == "" {
		if pt := rt; pt.Kind() == reflect.Ptr {
			star = "*"
			rt = pt
		}
	}
	if rt.Name() != "" {
		if rt.PkgPath() == "" {
			return star + rt.Name()
		} else {
			return star + rt.PkgPath() + "." + rt.Name()
		}
	}
	return name
}

func handler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	buf := &bytes.Buffer{}
	args := []interface{}{Object{}, &Object{}, submodule.Object{}, &submodule.Object{}, "", []submodule.Object{}}
	gob.NewEncoder(buf).Encode(args)
	c.Infof("payload=%s", buf.Bytes())

	for _, v := range args {
		t := reflect.ValueOf(v).Type()
		c.Infof("PkgPath=%s Name=%s String=%s Name()=%s %%T=%T", t.PkgPath(), t.Name(), t.String(), Name(v), v)
		matched, _ := regexp.MatchString(`main\d+`, reflect.TypeOf(v).PkgPath())
		// gob.RegisterName(reflect.TypeOf(v).String(), v)
		c.Infof("matched=%v", matched)
	}
}

var laterFunc = delay.Func("key", func(c appengine.Context, obj Object) {
	c.Infof("later")
})

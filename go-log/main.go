package hello

import (
	"net/http"
	"strings"

	"appengine"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	for i := 0; i < 20; i++ {
		c.Infof("%d %s", i, strings.Repeat("x", 8000))
	}
}

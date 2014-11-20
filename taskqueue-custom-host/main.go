package hello

import (
	"net/http"
	"net/url"

	"appengine"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	t := taskqueue.NewPOSTTask("/not-a-handler", url.Values{})
	err := taskqueue.Add(c, t, "")
	if err != nil {
		c.Errorf("err: %+v", err)
	}
}

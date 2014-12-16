package main

import (
	"net/http"

	"appengine"
)

func init() {
	http.HandleFunc("/_ah/start", func(w http.ResponseWriter, r *http.Request) {
		c := appengine.NewContext(r)
		c.Infof("start")
	})
}

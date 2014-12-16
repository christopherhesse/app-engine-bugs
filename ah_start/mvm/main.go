package main

import (
	"net/http"

	"golang.google.com/appengine"
	// "golang.google.com/appengine/internal"
)

func init() {
	http.HandleFunc("/_ah/start", func(w http.ResponseWriter, r *http.Request) {
		c := appengine.NewContext(r)
		// c := internal.BackgroundContext()
		c.Infof("start")
	})
}

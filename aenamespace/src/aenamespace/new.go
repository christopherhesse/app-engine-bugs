package main

import (
	"net/http"
	"net/url"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/memcache"
	"google.golang.org/appengine/taskqueue"
)

func init() {
	http.HandleFunc("/new-enqueue", newEnqueueHandler)
	http.HandleFunc("/new-task", newTaskHandler)
}

func newEnqueueHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	ctx, err := appengine.Namespace(ctx, "new-namespace")
	if err != nil {
		log.Infof(ctx, "error using namespace: %v", err)
		return
	}

	item := &memcache.Item{
		Key:   "new-key",
		Value: []byte("new-namespace"),
	}
	if err := memcache.Set(ctx, item); err != nil {
		log.Infof(ctx, "error setting item: %v", err)
		return
	}
	t := taskqueue.NewPOSTTask("/new-task", url.Values{})
	if _, err := taskqueue.Add(ctx, t, ""); err != nil {
		log.Infof(ctx, "error enqueuing task: %v", err)
		return
	}
}

func newTaskHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	// ctx, _ = appengine.Namespace(ctx, "new-namespace") // adding this line will fix the issue
	if item, err := memcache.Get(ctx, "new-key"); err == memcache.ErrCacheMiss {
		log.Infof(ctx, "item not in the cache")
	} else if err != nil {
		log.Infof(ctx, "error getting item: %v", err)
	} else {
		log.Infof(ctx, "the value is %q", item.Value)
	}
}

package main

import (
	"net/http"
	"net/url"

	"appengine"
	"appengine/memcache"
	"appengine/taskqueue"
)

func init() {
	http.HandleFunc("/classic-enqueue", classicEnqueueHandler)
	http.HandleFunc("/classic-task", classicTaskHandler)
}

func classicEnqueueHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	ctx, err := appengine.Namespace(ctx, "classic-namespace")
	if err != nil {
		ctx.Infof("error using namespace: %v", err)
		return
	}

	item := &memcache.Item{
		Key:   "classic-key",
		Value: []byte("classic-namespace"),
	}
	if err := memcache.Set(ctx, item); err != nil {
		ctx.Infof("error setting item: %v", err)
		return
	}
	t := taskqueue.NewPOSTTask("/classic-task", url.Values{})
	if _, err := taskqueue.Add(ctx, t, ""); err != nil {
		ctx.Infof("error enqueuing task: %v", err)
		return
	}
}

func classicTaskHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	if item, err := memcache.Get(ctx, "classic-key"); err == memcache.ErrCacheMiss {
		ctx.Infof("item not in the cache")
	} else if err != nil {
		ctx.Infof("error getting item: %v", err)
	} else {
		ctx.Infof("the value is %q", item.Value)
	}
}

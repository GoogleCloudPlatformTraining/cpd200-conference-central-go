package main

import (
	"net/http"
	"log"
	"appengine"
	"golang.org/appengine/mail"
)

func SetAnnouncementHandler(w http.ResponseWriter, r *http.Request) {
	//Set Announcement in Memcache.
	if r.Method != "GET" {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	header := r.Header.Get("X-AppEngine-Cron")
	if header == "" {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("attempt to access cron handler directly, missing custom App Engine header"))
		return
	}
	cacheAnnouncement(r)
	w.WriteHeader(http.StatusNoContent)
}

func init() {
	http.HandleFunc("/crons/set_announcement", SetAnnouncementHandler)
}

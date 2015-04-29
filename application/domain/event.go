package domain

import (
	"net/http"
	"time"
)

// An Event represents an interaction that could be of interest to external application to perform usage analytics.
type Event struct {
	Type string
	Time int64

	//Request-related information
	IP        string
	UserAgent string
	Referrer  string
	Protocol  string

	Link *Link
}

// NewLinkCreatedEvent generates an event of type "LinkCreated" for a supplied Link and http.Request
func NewLinkCreatedEvent(l *Link, r *http.Request) *Event {
	return newEvent("LinkCreated", l, r)
}

// NewLinkRetrievedEvent generates an event of type "LinkRetrieved" for a supplied Link and http.Request
func NewLinkRetrievedEvent(l *Link, r *http.Request) *Event {
	return newEvent("LinkRetrieved", l, r)
}

func newEvent(t string, l *Link, r *http.Request) *Event {
	return &Event{
		Type:      t,
		Time:      time.Now().Unix(),
		IP:        r.RemoteAddr,
		UserAgent: r.UserAgent(),
		Referrer:  r.Referer(),
		Link:      l,
	}
}

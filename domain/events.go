package domain

import (
    "time"
    "net/http"
)

type DomainEvent struct {
    Type string
    Time int64
    IP string
    UserAgent string
    Referrer string

    Link Link
}

type EventNotifier interface {
    Notify(event *DomainEvent)
}

func NewDomainEvent(t string, l Link, r *http.Request) *DomainEvent {
    return &DomainEvent{
        Type: t,
        Time: time.Now().Unix(),
        IP: r.RemoteAddr,
        UserAgent: r.UserAgent(),
        Referrer: r.Referer(),
        Link: l,
    }
}
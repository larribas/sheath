package domain

import (
    "net/url"

    "github.com/dchest/uniuri"
)

type Link struct {
    Original *url.URL
    Stub string
}

func NewLink(rawURL string) (Link, error) {
    // TODO Check URL validity constraints
    u, err := url.Parse(rawURL)
    if err != nil {
        return Link{}, err
    }

    return Link{Original: u, Stub: uniuri.New()}, nil
}

type LinkRepository interface {
    Store(Link) error
    Find(string) (Link, error)
}
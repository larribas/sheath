package main

import (
    // 3rd Party Libraries (try to reduce)
    "github.com/gorilla/pat"
    "github.com/dchest/uniuri"

    // Native Libraries
    "net/http"
    "errors"
    "net/url"
)

// DOMAIN

// Types
type Link struct {
    original url.URL
    shortened string
}

func NewLink(rawURL string) *Link {
    // TODO Check URL validity constraints
    // TODO Return error if it couldn't be parsed
    u, _ := url.Parse(rawURL)
    return &Link{
        original: *u,
        shortened: uniuri.New(),
    }
}


// Interfaces
type LinkRepository interface {
    Store(Link) error
    Find(string) (Link, error)
}


type InMemoryLinkRepository struct {
    // TODO Find a better name
    table map[string]url.URL
}

func (r InMemoryLinkRepository) Store(l Link) (error) {
    if _, exists := r.table[l.shortened]; exists {
        return errors.New("A link with such Hash already exists")
    } else {
        r.table[l.shortened] = l.original
        return nil
    }
}

func (r InMemoryLinkRepository) Find(hash string) (Link, error) {
    if url, exists := r.table[hash]; exists {
        return Link{original: url, shortened: hash}, nil
    } else {
        return Link{}, errors.New("There is no link with such Hash")
    }
}


// Protocol Layer: HTTP
var repo LinkRepository

func CreateLink(w http.ResponseWriter, r *http.Request) {
    // TODO Send this information to the Application Layer's corresponding command
    //fmt.Fprintf(w, "Host=%s, URL=%s, RemoteIP=%s, UserAgent=%s, Referrer=%s", r.Host, url, r.RemoteAddr, r.UserAgent(), r.Referer())

    url := r.PostFormValue("url")
    link := NewLink(url)
    if err := repo.Store(*link); err != nil {
        w.WriteHeader(http.StatusBadRequest)
    } else {
        w.WriteHeader(http.StatusCreated)
        w.Write([]byte(link.shortened))
    }
}

func RetrieveLink(w http.ResponseWriter, r *http.Request) {
    // TODO Make sense of the UserAgent. It includes everything!
    // TODO Send this information to the Application Layer's corresponding command
    // fmt.Fprintf(w, "Hash=%s, RemoteIP=%s, UserAgent=%s, Referrer=%s\n", hash, r.RemoteAddr, r.UserAgent(), r.Referer())

    // Actual logic
    hash := r.URL.Query().Get(":hash")
    if link, err := repo.Find(hash); err != nil {
        w.WriteHeader(http.StatusNotFound)
    } else {
        // TODO Compare with StatusMovedPermanently, StatusFound, and StatusSeeOther
        http.Redirect(w, r, link.original.String(), http.StatusTemporaryRedirect)
    }
}

func main() {
    // Initialize global variables
    // TODO Take initialization to constructor function
    repo = InMemoryLinkRepository{table: make(map[string]url.URL) }

    // TODO Get rid of Pat and try doing it myself
    // TODO Post pattern lets /anything pass through. Stop it
    r := pat.New()
    r.Get("/{hash}", RetrieveLink)
    r.Post("/", CreateLink)
    http.Handle("/", r)

    http.ListenAndServe(":1827", nil)
}
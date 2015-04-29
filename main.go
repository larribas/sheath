package main

import (
    "net/http"

    "github.com/bmizerany/pat"
    "github.com/medbook/sheathe/commands"
)


func CreateLink(w http.ResponseWriter, r *http.Request) {
    url := r.PostFormValue("url")
    if url == "" {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte("You must supply a `url` to shorten"))
        return
    }

    linkStub, err := commands.CreateLink(url, r)
    if err != nil {
        // TODO Distinguish "500 error" vs "400 error, bad or blacklisted URL"
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    w.WriteHeader(http.StatusCreated)
    w.Write([]byte(linkStub))
}

func RetrieveLink(w http.ResponseWriter, r *http.Request) {
    linkStub := r.URL.Query().Get(":short_id")

    originalLink, err := commands.RetrieveLink(linkStub, r)
    if err != nil {
        // TODO Distinguish "500 error" vs "404 error, URL not found"
        w.WriteHeader(http.StatusNotFound)
        return
    }

    // TODO Compare with StatusMovedPermanently, StatusFound, and StatusSeeOther
    http.Redirect(w, r, originalLink, http.StatusTemporaryRedirect)
    w.Write([]byte(linkStub))
}

func main() {
    // TODO Get rid of Pat and implement both endpoints myself
    // TODO Post pattern lets /anything pass through. Stop it
    r := pat.New()
    r.Get("/:short_id", http.HandlerFunc(RetrieveLink))
    r.Post("/", http.HandlerFunc(CreateLink))
    http.Handle("/", r)

    http.ListenAndServe(":1827", nil)
}
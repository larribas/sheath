package in_memory

import (
    "net/url"
    "errors"

    "github.com/medbook/sheathe/domain"
)


// In-memory LinkRepository implementation
type InMemoryLinkRepository struct {
    // TODO Find a better name
    table map[string]url.URL
}

func (r InMemoryLinkRepository) Store(l domain.Link) (error) {
    if _, exists := r.table[l.Stub]; exists {
        return errors.New("A link with such Hash already exists")
    } else {
        r.table[l.Stub] = l.Original
        return nil
    }
}

func (r InMemoryLinkRepository) Find(hash string) (domain.Link, error) {
    if url, exists := r.table[hash]; exists {
        return domain.Link{Original: url, Stub: hash}, nil
    } else {
        // TODO Domain Errors
        return domain.Link{}, errors.New("There is no link with such Hash")
    }
}

func NewInMemoryLinkRepository() *InMemoryLinkRepository {
    return &InMemoryLinkRepository{table: make(map[string]url.URL) }
}

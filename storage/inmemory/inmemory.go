// Package inmemory implements a LinkRepository that stores and retrieves links from the running application's heap.
// As such, changes won't be persistent. This implementation is intended for testing purposes
package inmemory

import (
	"net/url"

	"github.com/medbook/sheath/application/domain"
)

// LinkRepository holds a map of link stubs onto their original URLs,
// and allows for the storage and retrieval of links from that map
type LinkRepository struct {
	urlFor map[string]*url.URL
}

// NewLinkRepository returns a new instance of LinkRepository,
// ensuring it complies with the corresponding domain interface
func NewLinkRepository() domain.LinkRepository {
	return &LinkRepository{urlFor: make(map[string]*url.URL)}
}

// Store saves the specified Link to memory
func (r LinkRepository) Store(l *domain.Link) error {
	if _, exists := r.urlFor[l.Stub]; exists {
		return domain.ErrLinkAlreadyExists(l.Stub)
	}

	r.urlFor[l.Stub] = l.Original
	return nil
}

// Find retrieves the Link corresponding to the supplied stub from memory
func (r LinkRepository) Find(stub string) (*domain.Link, error) {
	if url, exists := r.urlFor[stub]; exists {
		return &domain.Link{Original: url, Stub: stub}, nil
	}

	return &domain.Link{}, domain.ErrLinkNotFound(stub)
}

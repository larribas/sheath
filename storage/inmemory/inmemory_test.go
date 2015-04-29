package inmemory

import (
	"testing"

	"github.com/larribas/sheath/test"
)

func TestLinkRepository(t *testing.T) {
	test.GenericLinkRepositoryTest(t, NewLinkRepository)
}

package inmemory

import (
	"testing"

	"github.com/medbook/sheath/test"
)

func TestLinkRepository(t *testing.T) {
	test.GenericLinkRepositoryTest(t, NewLinkRepository)
}

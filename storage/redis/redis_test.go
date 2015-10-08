package redis

import (
	"testing"

	"github.com/MedBrain/sheath/test"
)

func TestLinkRepository(t *testing.T) {
	test.GenericLinkRepositoryTest(t, NewLinkRepository)
}

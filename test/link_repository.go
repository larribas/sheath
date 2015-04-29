package test

import (
	"testing"

	"github.com/larribas/sheath/application/domain"
)

// GenericLinkRepositoryTest receives a testing type and a particular LinkRepository implementation, and submits it
// to a series of test that validate the implementation complies with the appropriate functionality
func GenericLinkRepositoryTest(t *testing.T, new func() domain.LinkRepository) {
	repo := new()

	testStoringALink(t, repo)
	testStoringALinkTwice(t, repo)
	testFindingALink(t, repo)
	testFindingANonexistentLink(t, repo)
}

func testStoringALink(t *testing.T, implementationUnderTest domain.LinkRepository) {
	link, _ := domain.NewLink("some.link")
	err := implementationUnderTest.Store(link)

	if err != nil {
		t.Error("Expected LinkRepository::Store to save a link successfully")
	}
}

func testStoringALinkTwice(t *testing.T, implementationUnderTest domain.LinkRepository) {
	link, _ := domain.NewLink("some.link")
	implementationUnderTest.Store(link)

	// We assume it stores it successfully the first time
	if _, ok := implementationUnderTest.Store(link).(domain.ErrLinkAlreadyExists); !ok {
		t.Error("Expected LinkRepository::Store to return 'ErrLinkAlreadyExists' when storing the same link twice")
	}
}

func testFindingALink(t *testing.T, implementationUnderTest domain.LinkRepository) {
	linkBefore, _ := domain.NewLink("some.link")
	implementationUnderTest.Store(linkBefore)

	// We assume it stores it successfully
	linkAfter, err := implementationUnderTest.Find(linkBefore.Stub)
	if err != nil {
		t.Error("Expected LinkRepository::find not to return an error when finding a valid link")
	}

	if linkAfter.Original.String() != linkBefore.Original.String() {
		t.Error("Expected LinkRepository::find to return the same link we previously saved")
	}
}

func testFindingANonexistentLink(t *testing.T, implementationUnderTest domain.LinkRepository) {
	_, err := implementationUnderTest.Find("nonexistent-link")
	if _, ok := err.(domain.ErrLinkNotFound); !ok {
		t.Error("Expected LinkRepository::find to return 'ErrLinkNotFound' when finding a nonexistent link")
	}
}

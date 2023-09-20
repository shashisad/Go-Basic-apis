package helpers

import (
	"github.com/onsi/gomega"
	"testing"
)

func TestFetchAllPosts_NoError(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	url1 := "https://jsonplaceholder.typicode.com/posts"
	_, err := FetchAllPosts(url1)
	g.Expect(err).ShouldNot(gomega.HaveOccurred())
}

func TestFetchAllPostsWith_UrlError(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	// Parsing wrong url host
	url := "https://jsonplaceholder.typicode.comm/posts"
	_, err := FetchAllPosts(url)
	g.Expect(err).ShouldNot(gomega.HaveOccurred())
}

func TestFetchAllPostsWith_UnMarshalError(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	url := "https://jsonplaceholder.typicode.com/comments/2"
	_, err := FetchAllPosts(url)
	g.Expect(err).Should(gomega.HaveOccurred())
}

func TestFetchCommentForPostID_NoError(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	_, err := FetchCommentForPostID(1)
	g.Expect(err).ShouldNot(gomega.HaveOccurred())
}

func TestFetchCommentForPostID_WrongIdError(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	_, err := FetchCommentForPostID(100000)
	g.Expect(err).Should(gomega.HaveOccurred())
}

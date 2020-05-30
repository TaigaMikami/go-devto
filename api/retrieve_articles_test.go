package api

import (
	"testing"
)

func TestRetrieveArticles(t *testing.T) {
	pageNum := 1
	perPageNum := 2
	client := NewClient("")
	opt := &RetrieveArticlesOption{
		Page:    pageNum,
		PerPage: perPageNum,
		State:   StateFresh,
	}
	articles, err := client.RetrieveArticles(opt)
	if err != nil {
		t.Fatal(err)
	}

	if len(articles) != perPageNum {
		t.Errorf("Got wrong number of articles expect: %d, actual: %d\n", perPageNum, len(articles))
	}
}

package devto

import "testing"

func TestRetrieveVideoArticles(t *testing.T) {
	pageNum := 1
	perPageNum := 2
	client := NewClient("")
	opt := &RetrieveVideoArticlesOption{
		Page:    pageNum,
		PerPage: perPageNum,
	}
	articles, err := client.RetrieveVideoArticles(opt)
	if err != nil {
		t.Fatal(err)
	}

	if len(articles) != perPageNum {
		t.Errorf("Got wrong number of articles expect: %d, actual: %d\n", perPageNum, len(articles))
	}
}

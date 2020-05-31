package devto

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

func TestRetrieveArticle(t *testing.T) {
	id := 150589
	client := NewClient("")
	article, err := client.RetrieveArticleById(id)
	if err != nil {
		t.Fatal(err)
	}

	if article.Id != id {
		t.Errorf("Got wrong id expect: %d, actual: %d\n", id, article.Id)
	}
}

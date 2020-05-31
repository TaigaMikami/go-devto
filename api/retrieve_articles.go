package api

import "strconv"

type RetrieveArticlesOption struct {
	Page         int
	PerPage      int
	Tag          string
	Username     string
	State        State
	Top          int
	CollectionId int
}

type State string

const (
	StateFresh  State = "fresh"
	StateRising       = "rising"
	StateAll          = "all"
)

func (c *Client) RetrieveArticles(option *RetrieveArticlesOption) ([]*Article, error) {
	var res []*Article
	err := GetWithQuery("/articles", option, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) RetrieveArticleById(id int) (*Article, error) {
	res := &Article{}
	err := SimpleGet("/articles/"+strconv.Itoa(id), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

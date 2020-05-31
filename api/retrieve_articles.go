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

type RetrieveUserArticlesOption struct {
	Page    int
	PerPage int
}

/* GET /articles */
func (c *Client) RetrieveArticles(option *RetrieveArticlesOption) ([]*Article, error) {
	var res []*Article
	err := GetWithQuery("/articles", c.ApiKey, option, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

/* GET /articles/:id */
func (c *Client) RetrieveArticleById(id int) (*Article, error) {
	res := &Article{}
	err := SimpleGet("/articles/"+strconv.Itoa(id), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

/* GET /articles/me */
func (c *Client) RetrieveUserArticles(option *RetrieveUserArticlesOption) ([]*Article, error) {
	var res []*Article
	err := GetWithQuery("/articles/me/published", c.ApiKey, option, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

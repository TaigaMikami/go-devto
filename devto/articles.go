package devto

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
	err := SimpleGet("/articles/"+strconv.Itoa(id), c.ApiKey, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

/* GET /articles/me */
func (c *Client) RetrieveUserArticles(option *RetrieveUserArticlesOption) ([]*Article, error) {
	var res []*Article
	err := GetWithQuery("/articles/me", c.ApiKey, option, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

/* GET /articles/me/published */
func (c *Client) RetrieveUserPublishedArticles(option *RetrieveUserArticlesOption) ([]*Article, error) {
	var res []*Article
	err := GetWithQuery("/articles/me/published", c.ApiKey, option, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

/* GET /articles/me/unpublished */
func (c *Client) RetrieveUserUnpublishedArticles(option *RetrieveUserArticlesOption) ([]*Article, error) {
	var res []*Article
	err := GetWithQuery("/articles/me/unpublished", c.ApiKey, option, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

/* GET /articles/me/all */
func (c *Client) RetrieveUserAllArticles(option *RetrieveUserArticlesOption) ([]*Article, error) {
	var res []*Article
	err := GetWithQuery("/articles/me/all", c.ApiKey, option, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type ArticleContent struct {
	Title          string   `json:"title"`
	BodyMarkdown   string   `json:"body_markdown"`
	Published      bool     `json:"published"`
	Series         string   `json:"series"`
	MainImage      string   `json:"main_image"`
	CanonicalUrl   string   `json:"canonical_url"`
	Description    string   `json:"description"`
	Tags           []string `json:"tags"`
	OrganizationId int      `json:"organization_id"`
}

type DraftArticle struct {
	Article ArticleContent `json:"article"`
}

/* POST /articles */
func (c *Client) AddArticle(option *DraftArticle) (*Article, error) {
	var res *Article
	err := PostWithJSON("/articles", c.ApiKey, option, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) ModifyArticle(id int, option *DraftArticle) (*Article, error) {
	var res *Article
	err := PutWithJSON("/articles/"+strconv.Itoa(id), c.ApiKey, option, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

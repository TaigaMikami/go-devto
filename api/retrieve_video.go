package api

type RetrieveVideoArticlesOption struct {
	Page    int
	PerPage int
}

/* GET /videos */
func (c *Client) RetrieveVideoArticles(option *RetrieveVideoArticlesOption) ([]*VideoArticle, error) {
	var res []*VideoArticle
	err := GetWithQuery("/articles", c.ApiKey, option, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

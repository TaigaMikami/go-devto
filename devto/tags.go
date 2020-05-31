package devto

type RetrieveTagsOption struct {
	Page    int
	PerPage int
}

/* GET /tags */
func (c *Client) RetrieveTags(option *RetrieveTagsOption) ([]*Tag, error) {
	var res []*Tag
	err := GetWithQuery("/tags", c.ApiKey, option, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

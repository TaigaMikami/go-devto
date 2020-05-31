package api

type RetrieveFollowersOption struct {
	Page    int
	PerPage int
}

/* GET /followers/users */
func (c *Client) RetrieveFollowers(option *RetrieveFollowersOption) ([]*Follower, error) {
	var res []*Follower
	err := GetWithQuery("/articles/me", c.ApiKey, option, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

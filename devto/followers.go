package devto

type RetrieveFollowersOption struct {
	Page    int
	PerPage int
}

/* GET /followers/users */
func (c *Client) RetrieveFollowers(option *RetrieveFollowersOption) ([]*Follower, error) {
	var res []*Follower
	err := GetWithQuery("/followers/users", c.ApiKey, option, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

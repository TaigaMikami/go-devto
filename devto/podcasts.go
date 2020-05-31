package devto

type RetrievePodcastOption struct {
	Page     int
	PerPage  int
	Username string
}

/* GET /podcast_episodes */
func (c *Client) RetrievePodcastEpisodes(option *RetrievePodcastOption) ([]*PodcastEpisode, error) {
	var res []*PodcastEpisode
	err := GetWithQuery("/podcast_episodes", c.ApiKey, option, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

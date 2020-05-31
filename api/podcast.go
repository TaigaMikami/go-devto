package api

type RetrievePodcast struct {
	Page     int
	PerPage  int
	Username string
}

/* PUT /podcast_episodes */
func (c *Client) RetrievePodcastEpisodes(option *RetrievePodcast) ([]*PodcastEpisode, error) {
	var res []*PodcastEpisode
	err := GetWithQuery("/podcast_episodes", c.ApiKey, option, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

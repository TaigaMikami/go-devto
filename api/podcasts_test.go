package api

import "testing"

func TestRetrievePodcastEpisodes(t *testing.T) {
	typeOf := "podcast_episodes"
	client := NewClient("")
	opt := &RetrievePodcastOption{
		Page:    1,
		PerPage: 1,
	}

	podcasts, err := client.RetrievePodcastEpisodes(opt)
	if err != nil {
		t.Fatal(err)
	}

	if podcasts[0].TypeOf != typeOf {
		t.Errorf("Got wrong type expect: %s, actual: %s\n", typeOf, podcasts[0].TypeOf)
	}
}

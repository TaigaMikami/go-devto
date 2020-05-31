package devto

import "testing"

func TestRetrieveTags(t *testing.T) {
	client := NewClient("")
	opt := &RetrieveTagsOption{
		Page: 1,
	}
	tags, err := client.RetrieveTags(opt)
	if err != nil {
		t.Fatal(err)
	}

	if tags[0].Name != "javascript" {
		t.Errorf("Got wrong name expect: %s, actual: %s\n", "javascritp", tags[0].Name)
	}
}

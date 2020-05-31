package devto

import (
	"testing"
)

func TestRetrieveComments(t *testing.T) {
	typeOf := "comment"
	client := NewClient("")
	opt := &RetrieveCommentsOption{
		AId: 270180,
	}
	comments, err := client.RetrieveComments(opt)
	if err != nil {
		t.Fatal(err)
	}

	if comments[0].TypeOf != typeOf {
		t.Errorf("Got wrong type expect: %s, actual: %s\n", typeOf, comments[0].TypeOf)
	}
}

func TestRetrieveCommentById(t *testing.T) {
	typeOf := "comment"
	id := "m51e"
	client := NewClient("")
	comment, err := client.RetrieveCommentById(id)
	if err != nil {
		t.Fatal(err)
	}

	if comment.TypeOf != typeOf {
		t.Errorf("Got wrong type expect: %s, actual: %s\n", typeOf, comment.TypeOf)
	}
}

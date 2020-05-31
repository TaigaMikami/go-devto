package api

type RetrieveCommentsOption struct {
	AId int
}

/* GET /comments */
func (c *Client) RetrieveComments(option *RetrieveCommentsOption) ([]*Comment, error) {
	var res []*Comment
	err := GetWithQuery("/comments", c.ApiKey, option, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

/* GET /comments/:id */
func (c *Client) RetrieveCommentById(id string) (*Comment, error) {
	res := &Comment{}
	err := SimpleGet("/comments/"+id, res)
	if err != nil {
		return nil, err
	}
	return res, err
}

package api

import "strconv"

/* GET /users/me */
func (c *Client) RetrieveUser(id int) (*User, error) {
	var res *User
	err := SimpleGet("/users/"+strconv.Itoa(id), c.ApiKey, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

/* GET /users/me */
func (c *Client) RetrieveAuthenticatedUser() (*User, error) {
	var res *User
	err := SimpleGet("/users/me", c.ApiKey, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

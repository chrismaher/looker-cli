package client

// GetUsers
func (c *client) GetUsers(email string, parameters map[string][]string) ([]Response, error) {
	ex, err := c.endpoint("/users/search/names", email)
	if err != nil {
		return nil, err
	}
	addParams(ex, parameters)

	return c.request(ex.String())
}

// GetRoles
func (c *client) GetRoles() ([]Response, error) {
	ex, err := c.endpoint("/roles")
	if err != nil {
		return nil, err
	}

	return c.request(ex.String())
}

//GetQueries
func (c *client) GetQueries() ([]Response, error) {
	ex, err := c.endpoint("/running_queries")
	if err != nil {
		return nil, err
	}

	return c.request(ex.String())
}

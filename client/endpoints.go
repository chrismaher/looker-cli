package client

// GetUsers
func (c *client) GetUsers(email string, parameters map[string][]string) []Response {
	ex := c.endpoint("/users/search/names", email)
	addParams(ex, parameters)

	return c.request(ex.String())
}

// GetRoles
func (c *client) GetRoles() []Response {
	ex := c.endpoint("/roles")

	return c.request(ex.String())
}

//GetQueries
func (c *client) GetQueries() []Response {
	ex := c.endpoint("/running_queries")

	return c.request(ex.String())
}

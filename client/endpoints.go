package client

// GetUsers
func (c *client) GetUsers(email string, parameters map[string][]string) ([]Response, error) {
	return c.request(endpoint{route: "/users/search/names", params: parameters}, email)
}

// GetRoles
func (c *client) GetRoles() ([]Response, error) {
	return c.request(endpoint{route: "/roles"})
}

//GetQueries
func (c *client) GetQueries() ([]Response, error) {
	return c.request(endpoint{route: "/running_queries"})
}

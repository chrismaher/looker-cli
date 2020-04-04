package client

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"
)

type client struct {
	*config
	token string
}

type endpoint struct {
	route  string
	params map[string][]string
}

// Response type will most often be used as []Response
type Response map[string]interface{}

// mkEndpoint takes a base endpoint, e.g. /running_queries and optional
// additional paths, and returns a *URL with a complete Looker API URL
func (c *client) mkEndpoint(base string, ps ...string) (*url.URL, error) {
	u, err := url.Parse(c.Path)
	if err != nil {
		return nil, err
	}
	u.Path = path.Join("/api/3.0", base)

	for _, p := range ps {
		u.Path = path.Join(u.Path, p)
	}

	return u, nil
}

// addParams uses a map to add an arbitrary number of URL parameters
// to an existing *URL value
func addParams(u *url.URL, ps map[string][]string) {
	params := url.Values{}
	for k, v := range ps {
		params.Add(k, strings.Join(v, ", "))
	}
	u.RawQuery = params.Encode()
}

// get makes GET HTTP requests to Looker API endpoints
func (c *client) get(urlstring string) ([]byte, error) {
	access := "token " + c.token
	req, _ := http.NewRequest("GET", urlstring, nil)
	req.Header.Add("Authorization", access)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(res.Body)
}

// request takes a URL string, calls get(), and unmarshalls the JSON Response
func (c *client) request(ep endpoint, ps ...string) ([]Response, error) {
	ex, err := c.mkEndpoint(ep.route, ps...)
	if err != nil {
		return nil, err
	}

	if ep.params != nil {
		addParams(ex, ep.params)
	}

	buf, err := c.get(ex.String())
	if err != nil {
		return nil, err
	}

	var list []Response
	if err := json.Unmarshal(buf, &list); err != nil {
		return nil, err
	}

	return list, nil
}

// getToken authenticates the client and sets client.token to the returned access_token
func (c *client) getToken() error {
	data := url.Values{}
	data.Set("client_id", c.ID)
	data.Add("client_secret", c.Secret)

	u, _ := url.Parse(c.Path)
	u.Path = "/login"
	urlStr := u.String()

	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, strings.NewReader(data.Encode()))

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var ret Response
	if err := json.Unmarshal(buf, &ret); err != nil {
		return err
	}

	c.token = ret["access_token"].(string)
	return nil
}

// New is a constructor for creating a client
func New() (*client, error) {
	c := &client{}
	config, err := readConfig()
	if err != nil {
		return nil, err
	}
	c.config = config

	if err := c.getToken(); err != nil {
		return nil, err
	}

	return c, nil
}

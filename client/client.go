package client

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
	"strings"
)

type client struct {
	*config
	token string
}

// Response type will most often be used as []Response
type Response map[string]interface{}

// endpoint takes a base endpoint, e.g. /running_queries and optional
// additional paths, and returns a *URL with a complete Looker API URL
func (c *client) endpoint(base string, ps ...string) *url.URL {
	u, err := url.Parse(c.Path)
	if err != nil {
		log.Fatal(err)
	}
	u.Path = path.Join("/api/3.0", base)

	for _, p := range ps {
		u.Path = path.Join(u.Path, p)
	}

	return u
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
func (c *client) get(urlstring string) []byte {
	access := "token " + c.token
	req, _ := http.NewRequest("GET", urlstring, nil)
	req.Header.Add("Authorization", access)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	buf, _ := ioutil.ReadAll(res.Body)

	return buf
}

// request takes a URL string, calls get(), and unmarshalls the JSON Response
func (c *client) request(ep string) (list []Response) {
	buf := c.get(ep)
	if err := json.Unmarshal(buf, &list); err != nil {
		log.Fatal(err)
	}

	return list
}

// getToken authenticates the client and sets client.token to the returned access_token
func (c *client) getToken() {
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
		log.Fatal(err)
	}

	buf, _ := ioutil.ReadAll(res.Body)

	var ret Response

	if err := json.Unmarshal(buf, &ret); err != nil {
		log.Fatal(err)
	}

	c.token = ret["access_token"].(string)
}

// New is a constructor for creating a client
func New() *client {
	c := &client{}
	c.config = readConfig()
	c.getToken()
	return c
}

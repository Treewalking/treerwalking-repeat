package jira

import (
	"io/ioutil"
	"net/http"
)

type JIRA struct {
	Client  *http.Client
	BaseURL string
}

func (api *JIRA) DoStuff() ([]byte, error) {
	resp, err := api.Client.Get(api.BaseURL + "/some/path")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	// handling error and doing stuff with body that needs to be unit tested
	return body, err
}

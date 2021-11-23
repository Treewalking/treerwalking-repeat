package jira_test

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"treerwalking.org/repeat/jira"
)

func TestDoStuffWithTestServer(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		assert.Equal(t, req.URL.String(), "/some/path")
		// Send response to be tested
		rw.Write([]byte(`OK`))
	}))
	// Close the server when test finishes
	defer server.Close()

	// Use Client & URL from our local test server
	api := jira.JIRA{server.Client(), server.URL}
	body, _ := api.DoStuff()

	assert.Equal(t, []byte("OK"), body)
}

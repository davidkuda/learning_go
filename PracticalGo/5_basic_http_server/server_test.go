// Two categories:
// 1. Integration Tests: Start the server and send requests to the endpoints.
// 2. Unit Test: Call the handlers with specified r and w objects.

package main

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		expected string
	}{
		{
			name:     "index",
			path:     "/v1/api",
			expected: "Hello World",
		},
		{
			name:     "healthcheck",
			path:     "/v1/healthz",
			expected: "Health shining like the moon!",
		},
	}

	mux := http.NewServeMux()
	setupHandlers(mux)

	ts := httptest.NewServer(mux)
	defer ts.Close()

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := http.Get(ts.URL + tc.path)
			if err != nil {
				log.Fatal(err)
			}
			respBody, err := io.ReadAll(resp.Body)
			resp.Body.Close()
			if err != nil {
				log.Fatal(err)
			}

			if string(respBody) != tc.expected {
				t.Errorf("Expected: %s, Got: %s", tc.expected, string(respBody))
			}
		})
	}
}

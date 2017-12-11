package profitbricks

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendRetryingRequest(t *testing.T) {
	cases := []struct {
		description     string
		responsePayload string
		retries         int
		requestPayload  string
		statusCode      int
		userAgent       string
	}{
		{
			description:     "Request without payload",
			responsePayload: "Request successful",
			retries:         2,
			requestPayload:  "",
			statusCode:      http.StatusOK,
			userAgent:       AgentHeader,
		},
		{
			description:     "Request with payload",
			responsePayload: "Request successful",
			retries:         2,
			requestPayload:  "Request Payload",
			statusCode:      http.StatusCreated,
			userAgent:       AgentHeader,
		},
	}

	for _, testCase := range cases {
		currentTry := 1
		requestsReceivedCount := 0
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestsReceivedCount = requestsReceivedCount + 1
			if testCase.requestPayload != "" {
				payload, err := ioutil.ReadAll(r.Body)
				assert.NoError(t, err, testCase.description)
				assert.Equal(t, testCase.requestPayload, string(payload), testCase.description)
			}

			assert.Equal(t, testCase.userAgent, r.Header.Get("User-Agent"))
			if currentTry <= testCase.retries {
				w.Header().Set("Retry-After", "1")
				w.WriteHeader(http.StatusTooManyRequests)
			} else {
				w.WriteHeader(testCase.statusCode)
				w.Write([]byte(testCase.responsePayload))
			}

			currentTry = currentTry + 1
		}))
		defer ts.Close()

		var request *http.Request
		if testCase.requestPayload == "" {
			request, _ = http.NewRequest("GET", ts.URL, nil)
		} else {
			request, _ = http.NewRequest("POST", ts.URL, bytes.NewBufferString(testCase.requestPayload))
		}

		resp, respErr := sendRetryingRequest(request)
		assert.NoError(t, respErr, testCase.description)
		payload, err := ioutil.ReadAll(resp.Body)
		assert.NoError(t, err, testCase.description)
		assert.Equal(t, testCase.responsePayload, string(payload), testCase.description)
		assert.Equal(t, testCase.retries+1, requestsReceivedCount, testCase.description)
		assert.Equal(t, testCase.statusCode, resp.StatusCode, testCase.description)
	}
}

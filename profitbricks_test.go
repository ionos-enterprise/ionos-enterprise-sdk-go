package profitbricks

import (
	"net/http"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/suite"
)

type SuiteClient struct {
	ClientBaseSuite
}

func Test_Client(t *testing.T) {
	suite.Run(t, new(SuiteClient))
}
func (s *SuiteClient) Test_Retry() {
	called := 0
	httpmock.RegisterResponder(http.MethodGet, `=~/?depth=10`,
		func(*http.Request) (*http.Response, error) {
			called++
			switch called {
			case 1:
				rsp := httpmock.NewBytesResponse(http.StatusTooManyRequests, []byte{})
				rsp.Header.Set("Retry-After", "1") // Overruled by RetryMaxWaitTime of 2 ns
				return rsp, nil
			case 2:
				return httpmock.NewBytesResponse(http.StatusBadGateway, []byte{}), nil
			case 3:
				return httpmock.NewBytesResponse(http.StatusGatewayTimeout, []byte{}), nil
			case 4:
				// Regression test for missing Retry-After in header
				rsp := httpmock.NewBytesResponse(http.StatusTooManyRequests, []byte{})
				return rsp, nil
			}
			// More response code
			return httpmock.NewBytesResponse(http.StatusOK, []byte{}), nil
		},
	)
	s.c.SetRetryCount(4)
	// slower that 2 nano seconds will result in a wait time of max int nano seconds (caused by internal normalization
	// in go resty
	s.c.SetRetryWaitTime(2 * time.Nanosecond)
	s.c.SetRetryMaxWaitTime(2 * time.Nanosecond)

	err := s.c.GetOK("/", nil)
	s.Error(err)
	s.Equal(4, called)
}

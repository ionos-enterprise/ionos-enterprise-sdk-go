package profitbricks

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/suite"
)

type SuiteToken struct {
	suite.Suite
	client *Client
	apiUrl string
}

func Test_SuiteToken(t *testing.T) {
	suite.Run(t, new(SuiteToken))
}

func (s *SuiteToken) SetupTest() {
	token := "eyJ0eXAiOiJKV1QiLCJraWQiOiJjMGY2MDQ4Yi1jZTg3LTRmOGEtODViMi01OTY3ZGI5YTA5NjEiLCJhbGciOiJSUzI1NiJ9.foo.bar"
	s.client = NewClientbyToken(token)
	s.apiUrl = s.client.client.authApiUrl
	httpmock.Activate()
}

func (s *SuiteToken) TearDownTest() {
	httpmock.DeactivateAndReset()
}

func (s *SuiteToken) Test_TokenID() {
	id, err := s.client.TokenID()
	if s.NoError(err) {
		s.Equal("c0f6048b-ce87-4f8a-85b2-5967db9a0961", id)
	}
}

func (s *SuiteToken) Test_DeleteCurrentToken() {
	query := url.Values{
		"criteria": []string{"CURRENT"},
	}
	httpmock.RegisterResponderWithQuery(http.MethodDelete, s.apiUrl+"/tokens", query,
		httpmock.NewStringResponder(200, ""))
	err := s.client.DeleteCurrentToken()
	s.NoError(err)
	s.Equal(1, httpmock.GetTotalCallCount())
}

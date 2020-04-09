package handler_test

import (
	"net/http"
	"testing"

	"github.com/dynastymasra/whistleblower/config"
	"github.com/dynastymasra/whistleblower/infrastructure/web/handler"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type PingSuite struct {
	suite.Suite
}

func Test_PingSuite(t *testing.T) {
	suite.Run(t, new(PingSuite))
}

func (p *PingSuite) SetupSuite() {
	config.SetupTestLogger()
}

func (p *PingSuite) Test_PingHandler() {
	assert.HTTPSuccess(p.T(), handler.Ping(), http.MethodGet, "/ping", nil)
	assert.HTTPBodyContains(p.T(), handler.Ping(), http.MethodGet, "/ping", nil, "{\"status\":\"success\"}")
}

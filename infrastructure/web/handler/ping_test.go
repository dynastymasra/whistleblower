package handler_test

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dynastymasra/cookbook"
	uuid "github.com/satori/go.uuid"

	"github.com/jinzhu/gorm"

	"github.com/dynastymasra/whistleblower/config"
	"github.com/dynastymasra/whistleblower/infrastructure/web/handler"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type PingSuite struct {
	suite.Suite
	db *gorm.DB
}

func Test_PingSuite(t *testing.T) {
	suite.Run(t, new(PingSuite))
}

func (p *PingSuite) SetupSuite() {
	config.Load()
	config.SetupTestLogger()
}

func (p *PingSuite) SetupTest() {
	db, err := config.Postgres().Client()
	if err != nil {
		log.Fatal(err)
	}

	p.db = db
}

func (p *PingSuite) Test_PingHandler() {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/ping", nil)
	ctx := context.WithValue(r.Context(), cookbook.RequestID, uuid.NewV4().String())

	handler.Ping(p.db)(w, r.WithContext(ctx))

	assert.Equal(p.T(), http.StatusOK, w.Code)
}

func (p *PingSuite) Test_PingHandler_Failed() {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/ping", nil)
	ctx := context.WithValue(r.Context(), cookbook.RequestID, uuid.NewV4().String())

	p.db.Close()
	handler.Ping(p.db)(w, r.WithContext(ctx))

	assert.Equal(p.T(), http.StatusInternalServerError, w.Code)
}

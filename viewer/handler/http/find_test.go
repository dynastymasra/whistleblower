package http_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	handler "github.com/dynastymasra/whistleblower/viewer/handler/http"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"

	"github.com/dynastymasra/cookbook"
	"github.com/dynastymasra/whistleblower/config"
	"github.com/dynastymasra/whistleblower/infrastructure/provider"
	"github.com/dynastymasra/whistleblower/viewer/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type StatisticCountSuite struct {
	suite.Suite
	viewerService *test.MockViewerService
}

func Test_StatisticCountSuite(t *testing.T) {
	suite.Run(t, new(StatisticCountSuite))
}

func (s *StatisticCountSuite) SetupSuite() {
	config.SetupTestLogger()
}

func (s *StatisticCountSuite) SetupTest() {
	s.viewerService = &test.MockViewerService{}
}

func (s *StatisticCountSuite) Test_StatisticCount_Success() {
	id := uuid.NewV4().String()

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/counter/v1/statistics/article_id/%s", id), nil)

	r = mux.SetURLVars(r, map[string]string{
		"article_id": id,
	})

	ctx := context.WithValue(r.Context(), cookbook.RequestID, uuid.NewV4().String())

	query := provider.NewQuery(config.ViewerTableName)
	query = query.Filter("article_id", provider.Equal, id)

	s.viewerService.On("Statistic", ctx, query).Return(&cookbook.JSON{}, nil)

	handler.StatisticCount(s.viewerService)(w, r.WithContext(ctx))

	assert.Equal(s.T(), http.StatusOK, w.Code)
}

func (s *StatisticCountSuite) Test_StatisticCount_Failed() {
	id := uuid.NewV4().String()

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/counter/v1/statistics/article_id/%s", id), nil)

	r = mux.SetURLVars(r, map[string]string{
		"article_id": id,
	})

	ctx := context.WithValue(r.Context(), cookbook.RequestID, uuid.NewV4().String())

	query := provider.NewQuery(config.ViewerTableName)
	query = query.Filter("article_id", provider.Equal, id)

	s.viewerService.On("Statistic", ctx, query).Return((*cookbook.JSON)(nil), assert.AnError)

	handler.StatisticCount(s.viewerService)(w, r.WithContext(ctx))

	assert.Equal(s.T(), http.StatusInternalServerError, w.Code)
}

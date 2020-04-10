package http_test

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dynastymasra/cookbook"
	"github.com/dynastymasra/whistleblower/config"
	handler "github.com/dynastymasra/whistleblower/viewer/handler/http"
	"github.com/dynastymasra/whistleblower/viewer/test"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CountViewerSuite struct {
	suite.Suite
	viewerRepo *test.MockViewerRepository
}

func Test_CreateArticleSuite(t *testing.T) {
	suite.Run(t, new(CountViewerSuite))
}

func (c *CountViewerSuite) SetupSuite() {
	config.SetupTestLogger()
}

func (c *CountViewerSuite) SetupTest() {
	c.viewerRepo = &test.MockViewerRepository{}
}

type errReader int

func (errReader) Read([]byte) (n int, err error) {
	return 0, assert.AnError
}

func viewerPayload(id string) []byte {
	return []byte(fmt.Sprintf(`{
		"data": {
			"id": "%s",
			"article_id": "%s"
		}
	}`, id, id))
}

//func (c *CountViewerSuite) Test_CountViewer_Success() {
//	id := uuid.NewV4().String()
//	payload := domain.Viewer{
//		ID:        id,
//		ArticleID: id,
//	}
//
//	w := httptest.NewRecorder()
//	r := httptest.NewRequest(http.MethodPost, "/counter/v1/statistics", bytes.NewReader(viewerPayload(id)))
//
//	ctx := context.WithValue(r.Context(), cookbook.RequestID, uuid.NewV4().String())
//
//	c.viewerRepo.On("Create", ctx, payload).Return(&domain.Article{}, nil)
//
//	handler.CountViewer(c.viewerRepo)(w, r.WithContext(ctx))
//	assert.Equal(c.T(), http.StatusOK, w.Code)
//}

func (c *CountViewerSuite) Test_CountViewer_Failed_ReadBody() {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/counter/v1/statistics", errReader(0))

	ctx := context.WithValue(r.Context(), cookbook.RequestID, uuid.NewV4().String())

	handler.CountViewer(c.viewerRepo)(w, r.WithContext(ctx))
	assert.Equal(c.T(), http.StatusBadRequest, w.Code)
}

func (c *CountViewerSuite) Test_CountViewer_Failed_Unmarshal() {
	reqInBytes := []byte(`<- test chan`)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/counter/v1/statistics", bytes.NewReader(reqInBytes))

	ctx := context.WithValue(r.Context(), cookbook.RequestID, uuid.NewV4().String())

	handler.CountViewer(c.viewerRepo)(w, r.WithContext(ctx))
	assert.Equal(c.T(), http.StatusBadRequest, w.Code)
}

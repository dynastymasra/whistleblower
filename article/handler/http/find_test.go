package http_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dynastymasra/cookbook"
	handler "github.com/dynastymasra/whistleblower/article/handler/http"
	"github.com/dynastymasra/whistleblower/domain"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"

	"github.com/dynastymasra/whistleblower/article/test"
	"github.com/dynastymasra/whistleblower/config"
	"github.com/stretchr/testify/suite"
)

type FindAllArticleSuite struct {
	suite.Suite
	articleRepo *test.MockArticleRepository
}

func Test_FindAllArticleSuite(t *testing.T) {
	suite.Run(t, new(FindAllArticleSuite))
}

func (f *FindAllArticleSuite) SetupSuite() {
	config.SetupTestLogger()
}

func (f *FindAllArticleSuite) SetupTest() {
	f.articleRepo = &test.MockArticleRepository{}
}

func (f *FindAllArticleSuite) Test_ProductFindAll_Success() {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/v1/articles", nil)

	ctx := context.WithValue(r.Context(), cookbook.RequestID, uuid.NewV4().String())

	f.articleRepo.On("FindAll", ctx, (map[string]interface{})(nil)).Return([]*domain.Article{{ID: "id"}}, nil)

	handler.FindAllArticle(f.articleRepo)(w, r.WithContext(ctx))

	assert.Equal(f.T(), http.StatusOK, w.Code)
}

func (f *FindAllArticleSuite) Test_ProductFindAll_Failed() {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/v1/articles", nil)

	ctx := context.WithValue(r.Context(), cookbook.RequestID, uuid.NewV4().String())

	f.articleRepo.On("FindAll", ctx, (map[string]interface{})(nil)).Return([]*domain.Article{}, assert.AnError)

	handler.FindAllArticle(f.articleRepo)(w, r.WithContext(ctx))

	assert.Equal(f.T(), http.StatusInternalServerError, w.Code)
}

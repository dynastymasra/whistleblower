package http_test

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dynastymasra/whistleblower/domain"

	"github.com/dynastymasra/cookbook"
	uuid "github.com/satori/go.uuid"

	handler "github.com/dynastymasra/whistleblower/article/handler/http"
	"github.com/dynastymasra/whistleblower/article/test"
	"github.com/dynastymasra/whistleblower/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CreateArticleSuite struct {
	suite.Suite
	articleService *test.MockArticleService
}

func Test_CreateArticleSuite(t *testing.T) {
	suite.Run(t, new(CreateArticleSuite))
}

func (c *CreateArticleSuite) SetupSuite() {
	config.SetupTestLogger()
}

func (c *CreateArticleSuite) SetupTest() {
	c.articleService = &test.MockArticleService{}
}

type errReader int

func (errReader) Read([]byte) (n int, err error) {
	return 0, assert.AnError
}

func articlePayload(id string) []byte {
	return []byte(fmt.Sprintf(`{
					"id": "%s"
					}`, id))
}

func (c *CreateArticleSuite) Test_CreateArticle_Success() {
	id := uuid.NewV4().String()
	payload := domain.Article{
		ID:        id,
		CreatedAt: nil,
		UpdatedAt: nil,
	}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/v1/articles", bytes.NewReader(articlePayload(id)))

	ctx := context.WithValue(r.Context(), cookbook.RequestID, uuid.NewV4().String())

	c.articleService.On("Create", ctx, payload).Return(&domain.Article{}, nil)

	handler.CreateArticle(c.articleService)(w, r.WithContext(ctx))
	assert.Equal(c.T(), http.StatusCreated, w.Code)
}

func (c *CreateArticleSuite) Test_CreateArticle_Failed_ReadBody() {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/v1/articles", errReader(0))

	ctx := context.WithValue(r.Context(), cookbook.RequestID, uuid.NewV4().String())

	handler.CreateArticle(c.articleService)(w, r.WithContext(ctx))
	assert.Equal(c.T(), http.StatusBadRequest, w.Code)
}

func (c *CreateArticleSuite) Test_CreateArticle_Failed_Unmarshal() {
	reqInBytes := []byte(`<- test chan`)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/v1/articles", bytes.NewReader(reqInBytes))

	ctx := context.WithValue(r.Context(), cookbook.RequestID, uuid.NewV4().String())

	handler.CreateArticle(c.articleService)(w, r.WithContext(ctx))
	assert.Equal(c.T(), http.StatusBadRequest, w.Code)
}

func (c *CreateArticleSuite) Test_CreateArticle_Failed() {
	id := uuid.NewV4().String()
	payload := domain.Article{
		ID:        id,
		CreatedAt: nil,
		UpdatedAt: nil,
	}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/v1/articles", bytes.NewReader(articlePayload(id)))

	ctx := context.WithValue(r.Context(), cookbook.RequestID, uuid.NewV4().String())

	c.articleService.On("Create", ctx, payload).Return((*domain.Article)(nil), assert.AnError)

	handler.CreateArticle(c.articleService)(w, r.WithContext(ctx))
	assert.Equal(c.T(), http.StatusInternalServerError, w.Code)
}

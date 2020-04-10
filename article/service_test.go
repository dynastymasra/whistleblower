package article_test

import (
	"context"
	"testing"

	"github.com/dynastymasra/whistleblower/domain"

	"github.com/stretchr/testify/assert"

	"github.com/dynastymasra/whistleblower/article"
	"github.com/dynastymasra/whistleblower/article/test"
	"github.com/dynastymasra/whistleblower/config"
	"github.com/stretchr/testify/suite"
)

type ServiceSuite struct {
	suite.Suite
	repo *test.MockArticleRepository
	*article.ServiceInstance
}

func Test_ServiceSuite(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}

func (s *ServiceSuite) SetupSuite() {
	config.SetupTestLogger()
}

func (s *ServiceSuite) SetupTest() {
	s.repo = &test.MockArticleRepository{}
	productService := article.NewService(s.repo)
	s.ServiceInstance = &productService
}

func (s *ServiceSuite) Test_Create_Success() {
	a := articleModel()
	s.repo.On("Create", context.Background(), a).Return(nil)
	s.repo.On("Find", context.Background(), map[string]interface{}{"id": a.ID}).Return(&domain.Article{}, nil)

	res, err := s.ServiceInstance.Create(context.Background(), a)

	assert.NotNil(s.T(), res)
	assert.NoError(s.T(), err)
}

func (s *ServiceSuite) Test_Create_Find_Failed() {
	a := articleModel()
	s.repo.On("Create", context.Background(), a).Return(nil)
	s.repo.On("Find", context.Background(), map[string]interface{}{"id": a.ID}).Return((*domain.Article)(nil), assert.AnError)

	res, err := s.ServiceInstance.Create(context.Background(), a)

	assert.NotNil(s.T(), res)
	assert.NoError(s.T(), err)
}

func (s *ServiceSuite) Test_Create_Failed() {
	a := articleModel()
	s.repo.On("Create", context.Background(), a).Return(assert.AnError)

	res, err := s.ServiceInstance.Create(context.Background(), a)

	assert.Nil(s.T(), res)
	assert.Error(s.T(), err)
}

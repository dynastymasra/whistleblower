package viewer_test

import (
	"context"
	"testing"
	"time"

	"github.com/dynastymasra/whistleblower/domain"
	"github.com/dynastymasra/whistleblower/infrastructure/provider"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"

	"github.com/dynastymasra/whistleblower/config"
	"github.com/dynastymasra/whistleblower/viewer"
	"github.com/dynastymasra/whistleblower/viewer/test"
	"github.com/stretchr/testify/suite"
)

type ServiceSuite struct {
	suite.Suite
	repo *test.MockViewerRepository
	*viewer.ServiceInstance
}

func Test_ServiceSuite(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}

func (s *ServiceSuite) SetupSuite() {
	config.SetupTestLogger()
}

func (s *ServiceSuite) SetupTest() {
	s.repo = &test.MockViewerRepository{}
	viewerService := viewer.NewService(s.repo)
	s.ServiceInstance = &viewerService
}

func result() []*domain.Viewer {
	articleID := uuid.NewV4().String()
	now := time.Now().UTC()
	fiveMinutesAgo := now.Add(time.Duration(-10) * time.Minute)
	aHourAgo := now.Add(time.Duration(-2) * time.Hour)
	aDayAgo := now.Add(time.Duration(-30) * time.Hour)
	twoDaysAgo := now.Add(time.Duration(-50) * time.Hour)
	threeDaysAgo := now.AddDate(0, 0, -4)
	return []*domain.Viewer{
		{
			ID:        uuid.NewV4().String(),
			ArticleID: articleID,
			CreatedAt: &fiveMinutesAgo,
		}, {
			ID:        uuid.NewV4().String(),
			ArticleID: articleID,
			CreatedAt: &aHourAgo,
		}, {
			ID:        uuid.NewV4().String(),
			ArticleID: articleID,
			CreatedAt: &aDayAgo,
		}, {
			ID:        uuid.NewV4().String(),
			ArticleID: articleID,
			CreatedAt: &twoDaysAgo,
		}, {
			ID:        uuid.NewV4().String(),
			ArticleID: articleID,
			CreatedAt: &threeDaysAgo,
		},
	}
}

func (s *ServiceSuite) Test_Statistic_Success() {
	s.repo.On("FindAll", context.Background()).Return(result(), nil)

	query := provider.NewQuery(config.ViewerTableName)
	res, err := s.ServiceInstance.Statistic(context.Background(), query)

	assert.NotNil(s.T(), res)
	assert.NoError(s.T(), err)
}

func (s *ServiceSuite) Test_Statistic_Failed() {
	s.repo.On("FindAll", context.Background()).Return([]*domain.Viewer{}, assert.AnError)

	query := provider.NewQuery(config.ViewerTableName)
	res, err := s.ServiceInstance.Statistic(context.Background(), query)

	assert.Nil(s.T(), res)
	assert.Error(s.T(), err)
}

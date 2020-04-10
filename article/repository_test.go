package article_test

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/dynastymasra/cookbook"

	"github.com/dynastymasra/whistleblower/domain"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"

	"github.com/dynastymasra/whistleblower/infrastructure/provider"

	"github.com/dynastymasra/whistleblower/config"

	"github.com/dynastymasra/whistleblower/article"
	"github.com/stretchr/testify/suite"
)

type RepositorySuite struct {
	suite.Suite
	*article.RepositoryInstance
}

func Test_RepositorySuite(t *testing.T) {
	suite.Run(t, new(RepositorySuite))
}

func (r *RepositorySuite) SetupSuite() {
	config.Load()
	config.SetupTestLogger()
}

func (r *RepositorySuite) TearDownSuite() {
	db, _ := config.Postgres().Client()
	provider.Close(db)
	provider.Reset()
}

func (r *RepositorySuite) SetupTest() {
	db, err := config.Postgres().Client()
	if err != nil {
		log.Fatal(err)
	}

	productRepo := article.NewRepository(db)

	r.RepositoryInstance = productRepo
}

func articleModel() domain.Article {
	timestamp := time.Now().UTC()
	return domain.Article{
		ID:        uuid.NewV4().String(),
		CreatedAt: &timestamp,
		UpdatedAt: &timestamp,
	}
}

func (r *RepositorySuite) Test_Create_Success() {
	err := r.RepositoryInstance.Create(context.Background(), articleModel())

	assert.NoError(r.T(), err)
}

func (r *RepositorySuite) Test_Create_Failed() {
	a := articleModel()
	a.ID = "test"

	err := r.RepositoryInstance.Create(context.Background(), a)

	assert.Error(r.T(), err)
}

func (r *RepositorySuite) Test_Find_Success() {
	a := articleModel()

	r.RepositoryInstance.Create(context.Background(), a)

	resp, err := r.RepositoryInstance.Find(context.Background(), cookbook.JSON{"id": a.ID})

	assert.NotNil(r.T(), resp)
	assert.NoError(r.T(), err)
}

func (r *RepositorySuite) Test_Find_Failed() {
	resp, err := r.RepositoryInstance.Find(context.Background(), cookbook.JSON{"id": uuid.NewV4().String()})

	assert.Nil(r.T(), resp)
	assert.Error(r.T(), err)
}

func (r *RepositorySuite) Test_FindAll_Success() {
	resp, err := r.RepositoryInstance.FindAll(context.Background(), cookbook.JSON{})

	assert.NotNil(r.T(), resp)
	assert.NoError(r.T(), err)
}

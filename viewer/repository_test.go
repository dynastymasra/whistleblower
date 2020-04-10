package viewer_test

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/dynastymasra/whistleblower/article"
	"github.com/dynastymasra/whistleblower/config"
	"github.com/dynastymasra/whistleblower/domain"
	"github.com/dynastymasra/whistleblower/infrastructure/provider"
	"github.com/dynastymasra/whistleblower/viewer"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/suite"
)

type RepositorySuite struct {
	suite.Suite
	viewerRepo  *viewer.RepositoryInstance
	articleRepo *article.RepositoryInstance
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

	viewerRepo := viewer.NewRepository(db)
	articleRepo := article.NewRepository(db)

	r.viewerRepo = viewerRepo
	r.articleRepo = articleRepo
}

func viewerModel(id string) domain.Viewer {
	timestamp := time.Now().UTC()
	return domain.Viewer{
		ID:        uuid.NewV4().String(),
		ArticleID: id,
		CreatedAt: &timestamp,
	}
}

func (r *RepositorySuite) Test_Create_Success() {
	articleModel := domain.Article{
		ID: uuid.NewV4().String(),
	}

	r.articleRepo.Create(context.Background(), articleModel)
	err := r.viewerRepo.Create(context.Background(), viewerModel(articleModel.ID))

	assert.NoError(r.T(), err)
}

func (r *RepositorySuite) Test_Create_Failed() {
	articleModel := domain.Article{
		ID: uuid.NewV4().String(),
	}

	r.articleRepo.Create(context.Background(), articleModel)
	err := r.viewerRepo.Create(context.Background(), viewerModel("test"))

	assert.Error(r.T(), err)
}

func (r *RepositorySuite) Test_FindAll_Success() {
	filter := provider.NewQuery(config.ViewerTableName)
	res, err := r.viewerRepo.FindAll(context.Background(), filter)

	assert.NoError(r.T(), err)
	assert.NotNil(r.T(), res)
}

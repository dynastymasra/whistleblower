package viewer

import (
	"context"
	"time"

	"github.com/dynastymasra/whistleblower/config"

	"github.com/dynastymasra/cookbook"
	"github.com/dynastymasra/whistleblower/infrastructure/provider"
)

var ()

type Service interface {
	Statistic(context.Context, *provider.Query) (*cookbook.JSON, error)
}

type ServiceInstance struct {
	repo Repository
}

func NewService(repo Repository) ServiceInstance {
	return ServiceInstance{repo: repo}
}

// Get counter article viewer max 5 days ago depend on business rules
// Count result by 5 minutes age, 1 hour ago, 1, 2, 3 days ago depend on business rules
// TODO: Improve the code for better implementation, in code or sql query
func (s ServiceInstance) Statistic(ctx context.Context, query *provider.Query) (*cookbook.JSON, error) {
	now := time.Now().UTC()
	fiveMinutesAgo := now.Add(time.Duration(-5) * time.Minute)
	aHourAgo := now.Add(time.Duration(-1) * time.Hour)
	aDayAgo := now.AddDate(0, 0, -1)
	twoDaysAgo := now.AddDate(0, 0, -2)
	threeDaysAgo := now.AddDate(0, 0, -3)
	fiveDaysAgo := now.AddDate(0, 0, -5)

	filter := query.Ordering(config.CreatedAtFieldName, provider.Descending)
	query.Filter(config.CreatedAtFieldName, provider.GreaterThanEqual, fiveDaysAgo)

	res, err := s.repo.FindAll(ctx, filter)
	if err != nil {
		return nil, err
	}

	// Initialize var counter
	countFiveMinutesAgo := 0
	countAHourAgo := 0
	countADayAgo := 0
	countTwoDaysAgo := 0
	countThreeDaysAgo := 0

	for _, viewer := range res {
		createdAt := viewer.CreatedAt
		if createdAt.After(aHourAgo) && createdAt.Before(fiveMinutesAgo) {
			countFiveMinutesAgo++
		} else if createdAt.After(aDayAgo) && createdAt.Before(aHourAgo) {
			countAHourAgo++
		} else if createdAt.After(twoDaysAgo) && createdAt.Before(aDayAgo) {
			countADayAgo++
		} else if createdAt.After(threeDaysAgo) && createdAt.Before(twoDaysAgo) {
			countTwoDaysAgo++
		} else if createdAt.Before(threeDaysAgo) {
			countThreeDaysAgo++
		}
	}

	count := []cookbook.JSON{
		{
			"reference": "5 minutes ago",
			"count":     countFiveMinutesAgo,
		}, {
			"reference": "1 hour ago",
			"count":     countAHourAgo,
		}, {
			"reference": "1 day ago",
			"count":     countADayAgo,
		}, {
			"reference": "2 days ago",
			"count":     countTwoDaysAgo,
		}, {
			"reference": "3 days ago",
			"count":     countThreeDaysAgo,
		},
	}

	return &cookbook.JSON{"count": count}, nil
}

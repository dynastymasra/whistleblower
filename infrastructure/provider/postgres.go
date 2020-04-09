package provider

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/matryer/resync"
	"github.com/sirupsen/logrus"
)

var (
	db      *gorm.DB
	err     error
	runOnce resync.Once
)

type Postgres struct {
	DatabaseName string
	Address      string
	Username     string
	Password     string
	MaxIdleConn  int
	MaxOpenConn  int
	LogEnabled   bool
}

func (p Postgres) Client() (*gorm.DB, error) {
	url := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", p.Username, p.Password, p.Address, p.DatabaseName)

	runOnce.Do(func() {
		db, err = gorm.Open("postgres", url)
		if err != nil {
			logrus.WithField("url", url).WithError(err).Errorln("Failed connect to database")
			return
		}

		db.DB().SetMaxIdleConns(p.MaxIdleConn)
		db.DB().SetMaxOpenConns(p.MaxOpenConn)
		db.LogMode(p.LogEnabled)
	})

	return db, err
}

package provider

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/matryer/resync"
	"github.com/sirupsen/logrus"

	_ "github.com/jinzhu/gorm/dialects/postgres"
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

func Ping(db *gorm.DB) error {
	if db == nil {
		return errors.New("does't have database connection")
	}
	return db.DB().Ping()
}

func Close(db *gorm.DB) error {
	if db == nil {
		return errors.New("does't have database data")
	}
	return db.DB().Close()
}

func Reset() {
	runOnce.Reset()
}

type Query struct {
	Model     string
	Limit     int
	Offset    int
	Filters   []*Filter
	Orderings []*Ordering
}

type Filter struct {
	Condition string
	Field     string
	Value     interface{}
}

type Ordering struct {
	Field     string
	Direction string
}

const (
	Equal            = "Equal"
	LessThan         = "LessThan"
	GreaterThan      = "GreaterThan"
	GreaterThanEqual = "GreaterThanEqual"
	LessThanEqual    = "LessThanEqual"
	JSON             = "JSON"

	Descending = "Descending"
	Ascending  = "Ascending"
)

var (
	validOrdering = map[string]bool{
		Descending: true,
		Ascending:  true,
	}
)

func NewQuery(model string) *Query {
	return &Query{
		Model: model,
	}
}

// Filter adds a filter to the query
func (q *Query) Filter(property, condition string, value interface{}) *Query {
	filter := NewFilter(property, condition, value)
	q.Filters = append(q.Filters, filter)
	return q
}

// Order adds a sort order to the query
func (q *Query) Ordering(property, direction string) *Query {
	order := NewOrdering(property, direction)
	q.Orderings = append(q.Orderings, order)
	return q
}

func (q *Query) Slice(offset, limit int) *Query {
	q.Offset = offset
	q.Limit = limit

	return q
}

// NewFilter creates a new property filter
func NewFilter(field, condition string, value interface{}) *Filter {
	return &Filter{
		Field:     field,
		Condition: condition,
		Value:     value,
	}
}

func NewOrdering(field, direction string) *Ordering {
	d := direction

	if !validOrdering[direction] {
		d = Descending
	}

	return &Ordering{
		Field:     field,
		Direction: d,
	}
}

func TranslateQuery(db *gorm.DB, query *Query) *gorm.DB {
	for _, filter := range query.Filters {
		switch filter.Condition {
		case Equal:
			q := fmt.Sprintf("%s = ?", filter.Field)
			db = db.Where(q, filter.Value)
		case GreaterThan:
			q := fmt.Sprintf("%s > ?", filter.Field)
			db = db.Where(q, filter.Value)
		case GreaterThanEqual:
			q := fmt.Sprintf("%s >= ?", filter.Field)
			db = db.Where(q, filter.Value)
		case LessThan:
			q := fmt.Sprintf("%s < ?", filter.Field)
			db = db.Where(q, filter.Value)
		case LessThanEqual:
			q := fmt.Sprintf("%s <= ?", filter.Field)
			db = db.Where(q, filter.Value)
		case JSON:
			q := fmt.Sprintf("%s @> ?", filter.Field)
			db = db.Where(q, filter.Value)
		default:
			q := fmt.Sprintf("%s = ?", filter.Field)
			db = db.Where(q, filter.Value)
		}
	}

	for _, order := range query.Orderings {
		switch order.Direction {
		case Ascending:
			o := fmt.Sprintf("%s %s", order.Field, "ASC")
			db = db.Order(o)
		case Descending:
			o := fmt.Sprintf("%s %s", order.Field, "DESC")
			db = db.Order(o)
		default:
			o := fmt.Sprintf("%s %s", order.Field, "DESC")
			db = db.Order(o)
		}
	}

	if query.Offset > 0 {
		db = db.Offset(query.Offset)
	}

	if query.Limit > 0 {
		db = db.Limit(query.Limit)
	}

	return db
}

package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gopkg.in/gormigrate.v1"
)

// database adapter
// should implement models/

type PostgresConfigStore struct {
	db *gorm.DB
}

func NewPostgresConfigStore(connect string) (*PostgresConfigStore, error) {

	db, err := gorm.Open("postgres", connect)
	if err != nil {
		return nil, err
	}

	if err = db.DB().Ping(); err != nil {
		return nil, err
	}

	return &PostgresConfigStore{
		db: db.Debug(),
	}, nil
}

func (s *PostgresConfigStore) Migrate() error {

	m := gormigrate.New(s.db, gormigrate.DefaultOptions, Migrations)

	return m.Migrate()
}

func (s *PostgresConfigStore) RollbackLast() error {

	m := gormigrate.New(s.db, gormigrate.DefaultOptions, Migrations)

	return m.RollbackLast()
}

func (s *PostgresConfigStore) Get(Type, ID string, result interface{}) error {
	err := s.db.Where("data = ?", ID).First(result).Error
	return err
}

func (s *PostgresConfigStore) Set(Type, ID string, model interface{}) error {
	err := s.db.Find(&model, ID).Error
	return err
}

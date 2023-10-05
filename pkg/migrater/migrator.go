package migrater

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type config interface {
	GetPostgresDB() string
	GetPostgresHost() string
	GetPostgresPass() string
	GetPostgresPort() string
	GetPostgresUser() string
	GetPostgresSSLMode() string
}

type migrater struct {
	user     string
	password string
	host     string
	port     string
	db       string
	sslMode  string
}

func New(cfg config) *migrater {
	return &migrater{
		user:     cfg.GetPostgresUser(),
		password: cfg.GetPostgresPass(),
		host:     cfg.GetPostgresHost(),
		port:     cfg.GetPostgresPort(),
		db:       cfg.GetPostgresDB(),
		sslMode:  cfg.GetPostgresSSLMode(),
	}
}

func (mig migrater) Migrate() error {
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", mig.user, mig.password, mig.host, mig.port, mig.db, mig.sslMode)
	fmt.Println(url)
	m, err := migrate.New("file://migrates", url)
	if err != nil {
		return err
	}
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}

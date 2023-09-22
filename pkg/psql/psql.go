package psql

import (
	"fmt"

	"github.com/go-pg/pg"
)

type configer interface {
	GetPostgresDB() string
	GetPostgresHost() string
	GetPostgresPass() string
	GetPostgresPort() string
	GetPostgresUser() string
	GetPostgresSSLMode() string
}

type postgresClient struct {
	host    string
	port    string
	user    string
	pass    string
	db      string
	sslMode string
}

func NewPostgresClient(cfg configer) *postgresClient {
	return &postgresClient{
		host:    cfg.GetPostgresHost(),
		port:    cfg.GetPostgresPort(),
		user:    cfg.GetPostgresUser(),
		pass:    cfg.GetPostgresPass(),
		db:      cfg.GetPostgresDB(),
		sslMode: cfg.GetPostgresSSLMode(),
	}
}

func (cl *postgresClient) GetDb() *pg.DB {
	return pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%s", cl.host, cl.port),
		User:     cl.user,
		Password: cl.pass,
		Database: cl.db,
	})
}

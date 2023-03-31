package dbrepo

import (
	"database/sql"

	"github.com/thiago-martinss/bookings/internal/config"
	"github.com/thiago-martinss/bookings/internal/repository"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}
}

package dbrepo

import (
	"database/sql"

	"github.com/alifarahani1998/bookings/controllers/config"
	"github.com/alifarahani1998/bookings/controllers/repository"
)


type postgresDBRepo struct {
	App *config.AppConfig
	DB *sql.DB
}




func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		DB: conn,
	}
}

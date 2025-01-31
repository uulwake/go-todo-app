package database

import (
	"database/sql"
	"os"
	"todo-app/internal/config"

	_ "github.com/lib/pq"
)

func NewPg(env *config.Env) (*sql.DB, error) {
	pg, err := sql.Open("postgres", env.PgUrl)
	if err != nil {
		return nil, err
	}

	pg.SetMaxOpenConns(10)
	pg.SetMaxIdleConns(2)

	if (os.Getenv("GO_ENV") == "production") {
		err = pg.Ping()
		if err != nil {
			return nil, err
		}
	}


	return pg, nil;
}
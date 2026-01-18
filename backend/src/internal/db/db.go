package db

import (
	"backend/db/minitec_db"
	"database/sql"
	"log/slog"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func New(url string) (*minitec_db.Queries, error) {
	db, err := sql.Open("mysql", url)
	if err != nil {
		slog.Error("Error connecting to database")
		return nil, err
	}
	
    driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		slog.Error("Error creating database instance")
		return nil, err
	}

    m, err := migrate.NewWithDatabaseInstance(
        "file://db/migrations",
        "mysql", 
        driver,
    )
	if err != nil {
		slog.Error("Error creating migration instance")
		return nil, err
	}
    
    err = m.Up()
	if err != nil {
		slog.Error("Error running migrations")
		return nil, err
	}

	queries := minitec_db.New(db)

	return queries, nil

}
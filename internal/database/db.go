package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/libsql/libsql-client-go/libsql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sahidrahman404/gigachad-api/ent"
)

var (
	defaultTimeout    = 3 * time.Second
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

type DB struct {
	Db  *sql.DB
	Ent *ent.Client
}

func New(dsn string) (*DB, error) {
	db, err := sql.Open("libsql", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(2 * time.Hour)

	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	drv := entsql.OpenDB(dialect.SQLite, db)
	entClient := ent.NewClient(ent.Driver(drv))
	return &DB{Db: db, Ent: entClient}, nil
}

func parseSqliteError(s string) string {
	if len(strings.Split(s, ":")) == 6 {
		return strings.TrimSpace(strings.Join(strings.Split(s, ":")[3:], ":"))
	}
	return s
}

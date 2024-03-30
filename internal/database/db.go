package database

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"time"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sahidrahman404/gigachad-api/ent"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
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

func New(dsn string, bucket string, s3Client *s3.Client) (*DB, error) {
	db, err := sql.Open("libsql", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxIdleTime(1 * time.Minute)
	db.SetConnMaxLifetime(2 * time.Hour)

	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	drv := entsql.OpenDB(dialect.SQLite, db)
	entClient := ent.NewClient(ent.Driver(drv), ent.DeleteObjectInput(&s3.DeleteObjectInput{Bucket: &bucket}), ent.S3Client(s3Client))
	return &DB{Db: db, Ent: entClient}, nil
}

func parseSqliteError(s string) string {
	if len(strings.Split(s, "\n")) == 2 {
		return strings.Split(s, "\n")[1]
	}
	return s
}

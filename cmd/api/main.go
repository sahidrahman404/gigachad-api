package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/debug"
	"sync"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/microcosm-cc/bluemonday"
	"github.com/sahidrahman404/gigachad-api/ent"
	_ "github.com/sahidrahman404/gigachad-api/ent/runtime"
	"github.com/sahidrahman404/gigachad-api/internal/aws"
	"github.com/sahidrahman404/gigachad-api/internal/database"
	"github.com/sahidrahman404/gigachad-api/internal/env"
	"github.com/sahidrahman404/gigachad-api/internal/img"
	"github.com/sahidrahman404/gigachad-api/internal/leveledlog"
	"github.com/sahidrahman404/gigachad-api/internal/purifier"
	"github.com/sahidrahman404/gigachad-api/internal/smtp"
	"github.com/sahidrahman404/gigachad-api/internal/version"
)

func main() {
	logger := leveledlog.NewLogger(os.Stdout, leveledlog.LevelAll, true)

	err := run(logger)
	if err != nil {
		trace := debug.Stack()
		logger.Fatal(err, trace)
	}
}

type config struct {
	baseURL  string
	httpPort int
	db       struct {
		dsn string
	}
	smtp struct {
		port     int
		host     string
		username string
		password string
		from     string
	}
	redis struct {
		address  string
		password string
	}
	img.Imgproxy
	aws.AWSConfig
	env string
}

type application struct {
	config        config
	logger        *leveledlog.Logger
	mailer        *smtp.Mailer
	wg            sync.WaitGroup
	storage       *database.Storage
	ent           *ent.Client
	presignClient *s3.PresignClient
	purifier      *bluemonday.Policy
}

func run(logger *leveledlog.Logger) error {
	var cfg config

	cfg.baseURL = env.GetString("BASE_URL", "http://localhost:4444")
	cfg.httpPort = env.GetInt("HTTP_PORT", 4444)
	cfg.db.dsn = env.GetString("DB_DSN", "db.sqlite")
	cfg.smtp.host = env.GetString("SMTP_HOST", "example.smtp.host")
	cfg.smtp.port = env.GetInt("SMTP_PORT", 25)
	cfg.smtp.username = env.GetString("SMTP_USERNAME", "example_username")
	cfg.smtp.password = env.GetString("SMTP_PASSWORD", "pa55word")
	cfg.smtp.from = env.GetString("SMTP_FROM", "Example Name <no_reply@example.org>")
	cfg.Key = env.GetString("KEY", "")
	cfg.Salt = env.GetString("SALT", "")
	cfg.ImgproxyHost = env.GetString("IMGPROXY_HOST", "")
	cfg.AccessKeyID = env.GetString("ACCESS_KEY_ID", "")
	cfg.SecretAccessKey = env.GetString("SECRET_ACCESS_KEY", "")
	cfg.OriginCDN = env.GetString("ORIGIN_CDN", "")
	cfg.AWSBucket = env.GetString("AWS_BUCKET", "")
	cfg.AWSRegion = env.GetString("AWS_REGION", "")
	cfg.env = env.GetString("ENV", "DEV")
	cfg.redis.address = env.GetString("REDIS_ADDRESS", "localhost:6379")
	cfg.redis.password = env.GetString("REDIS_PASSWORD", "")

	showVersion := flag.Bool("version", false, "display version and exit")

	flag.Parse()

	if *showVersion {
		fmt.Printf("version: %s\n", version.Get())
		return nil
	}

	db, err := database.New(cfg.db.dsn)
	if err != nil {
		return err
	}
	defer db.Db.Close()

	awsConfig, err := aws.NewAWSConfig(cfg.AWSConfig)
	if err != nil {
		return err
	}

	mailer := smtp.NewMailer(
		cfg.smtp.host,
		cfg.smtp.port,
		cfg.smtp.username,
		cfg.smtp.password,
		cfg.smtp.from,
	)

	p := purifier.NewPurifierPolicy()

	app := &application{
		config:        cfg,
		logger:        logger,
		mailer:        mailer,
		storage:       database.NewStorage(db.Ent),
		ent:           db.Ent,
		presignClient: aws.NewPresignClient(awsConfig),
		purifier:      p,
	}

	return app.serveHTTP()
}

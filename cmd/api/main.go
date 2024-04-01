package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"runtime/debug"
	"sync"

	"buf.build/gen/go/sahidrahman/gigachadapis/connectrpc/go/gigachad/v1/gigachadv1connect"
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
	env                 string
	reminderServiceHost string
}

type application struct {
	config                config
	logger                *leveledlog.Logger
	mailer                *smtp.Mailer
	wg                    sync.WaitGroup
	storage               *database.Storage
	ent                   *ent.Client
	presignClient         *s3.PresignClient
	purifier              *bluemonday.Policy
	s3Client              *s3.Client
	reminderServiceClient *gigachadv1connect.ReminderServiceClient
}

func run(logger *leveledlog.Logger) error {
	var cfg config

	cfg.baseURL = env.GetString("BASE_URL", "http://localhost:4444")
	cfg.httpPort = env.GetInt("HTTP_PORT", 4444)
	cfg.db.dsn = env.GetString("DB_DSN", "file:local.db")
	cfg.smtp.host = env.GetString("SMTP_HOST", "localhost")
	cfg.smtp.port = env.GetInt("SMTP_PORT", 1025)
	cfg.smtp.username = env.GetString("SMTP_USERNAME", "gigachad")
	cfg.smtp.password = env.GetString("SMTP_PASSWORD", "pa55word")
	cfg.smtp.from = env.GetString("SMTP_FROM", "Gigachad <hello@gigachad.buzz>")
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
	cfg.reminderServiceHost = env.GetString("REMINDER_SERVICE_ADDRESS", "http://localhost:8080")

	showVersion := flag.Bool("version", false, "display version and exit")

	flag.Parse()

	if *showVersion {
		fmt.Printf("version: %s\n", version.Get())
		return nil
	}

	awsConfig, err := aws.NewAWSConfig(cfg.AWSConfig)
	if err != nil {
		return err
	}

	s3Client := aws.NewS3Client(awsConfig)

	db, err := database.New(cfg.db.dsn, cfg.AWSBucket, s3Client)
	if err != nil {
		return err
	}
	defer db.Db.Close()

	mailer := smtp.NewMailer(
		cfg.smtp.host,
		cfg.smtp.port,
		cfg.smtp.username,
		cfg.smtp.password,
		cfg.smtp.from,
	)

	p := purifier.NewPurifierPolicy()

	r := gigachadv1connect.NewReminderServiceClient(http.DefaultClient, cfg.reminderServiceHost)

	app := &application{
		config:                cfg,
		logger:                logger,
		mailer:                mailer,
		storage:               database.NewStorage(db.Ent),
		ent:                   db.Ent,
		presignClient:         aws.NewPresignClient(awsConfig),
		purifier:              p,
		s3Client:              s3Client,
		reminderServiceClient: &r,
	}

	return app.serveHTTP()
}

package gql

import (
	"sync"

	"buf.build/gen/go/sahidrahman/gigachadapis/connectrpc/go/gigachad/v1/gigachadv1connect"
	"github.com/99designs/gqlgen/graphql"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/microcosm-cc/bluemonday"
	"github.com/sahidrahman404/gigachad-api"
	"github.com/sahidrahman404/gigachad-api/ent"
	"github.com/sahidrahman404/gigachad-api/internal/aws"
	"github.com/sahidrahman404/gigachad-api/internal/database"
	"github.com/sahidrahman404/gigachad-api/internal/img"
	"github.com/sahidrahman404/gigachad-api/internal/leveledlog"
	"github.com/sahidrahman404/gigachad-api/internal/smtp"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	client                *ent.Client
	mailer                *smtp.Mailer
	storage               *database.Storage
	logger                *leveledlog.Logger
	wg                    *sync.WaitGroup
	imgproxy              *img.Imgproxy
	awsCfg                *aws.AWSConfig
	purifier              *bluemonday.Policy
	s3Client              *s3.Client
	reminderServiceClient *gigachadv1connect.ReminderServiceClient
}

// NewSchema creates a graphql executable schema.
func NewSchema(
	c *ent.Client,
	m *smtp.Mailer,
	s *database.Storage,
	l *leveledlog.Logger,
	wg *sync.WaitGroup,
	i *img.Imgproxy,
	a *aws.AWSConfig,
	h *bluemonday.Policy,
	s3 *s3.Client,
	r *gigachadv1connect.ReminderServiceClient,
) graphql.ExecutableSchema {
	return gigachad.NewExecutableSchema(gigachad.Config{
		Resolvers: &Resolver{
			client:                c,
			mailer:                m,
			storage:               s,
			logger:                l,
			wg:                    wg,
			imgproxy:              i,
			awsCfg:                a,
			purifier:              h,
			s3Client:              s3,
			reminderServiceClient: r,
		},
	})
}

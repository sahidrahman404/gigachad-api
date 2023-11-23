package gql

import (
	"sync"

	"github.com/99designs/gqlgen/graphql"
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
	client   *ent.Client
	mailer   *smtp.Mailer
	storage  *database.Storage
	logger   *leveledlog.Logger
	wg       *sync.WaitGroup
	imgproxy *img.Imgproxy
	awsCfg   *aws.AWSConfig
	purifier *bluemonday.Policy
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
) graphql.ExecutableSchema {
	return gigachad.NewExecutableSchema(gigachad.Config{
		Resolvers: &Resolver{
			client:   c,
			mailer:   m,
			storage:  s,
			logger:   l,
			wg:       wg,
			imgproxy: i,
			awsCfg:   a,
			purifier: h,
		},
	})
}

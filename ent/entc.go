//go:build ignore

package main

import (
	"log"

	"buf.build/gen/go/sahidrahman/gigachadapis/connectrpc/go/gigachad/v1/gigachadv1connect"
	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var ReminderServiceClient gigachadv1connect.ReminderServiceClient

func main() {
	ex, err := entgql.NewExtension(
		// Tell Ent to generate a GraphQL schema for
		// the Ent schema in a file named ent.graphql.
		entgql.WithWhereInputs(true),
		entgql.WithConfigPath("./gqlgen.yml"),
		entgql.WithSchemaGenerator(),
		entgql.WithSchemaPath("./internal/gql/ent.graphql"),
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}
	opts := []entc.Option{
		entc.Dependency(
			entc.DependencyName("DeleteObjectInput"),
			entc.DependencyType(&s3.DeleteObjectInput{}),
		),
		entc.Dependency(
			entc.DependencyName("S3Client"),
			entc.DependencyType(&s3.Client{}),
		),
		entc.Extensions(ex),
	}
	if err := entc.Generate("./ent/schema", &gen.Config{
		Features: []gen.Feature{
			gen.FeaturePrivacy,
			gen.FeatureEntQL,
			gen.FeatureUpsert,
		},
	}, opts...); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}

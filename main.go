package main

import (
	"context"
	"log"

	"cloud.google.com/go/datastore"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/KazukiHayase/server-template/config"
	"github.com/KazukiHayase/server-template/graph/generated"
	"github.com/gin-contrib/cors"

	"github.com/KazukiHayase/server-template/graph/resolver"
	"github.com/gin-gonic/gin"
)

func main() {
	config, err := config.Environ()
	if err != nil {
		log.Fatalf("環境変数の設定に失敗しました: %v\n", err)
	}

	ctx := context.Background()
	dsClient, err := datastore.NewClient(ctx, config.GCP.ProjectID)
	if err != nil {
		log.Fatalf("datastoreの初期化に失敗しました: %v\n", err)
	}

	resolver := resolver.NewResolver(config, *dsClient)
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(
		generated.Config{Resolvers: &resolver},
	))

	r := gin.Default()

	corsConfig := cors.DefaultConfig()
	// TODO: 本番のドメイン追加
	corsConfig.AllowOrigins = []string{"http://localhost:3333"}
	r.Use(cors.New(corsConfig))

	r.POST("/graphql", func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	})

	if config.IsLocal() {
		r.GET("/", playgroundHandler())
	}

	r.Run(":8080")
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/graphql")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

package web

import (
	"context"
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/deployport/pavement/logging"
	"github.com/gitpushy/pave/cmd/pave/templates/pkg/configurations"
	"github.com/gitpushy/pave/cmd/pave/templates/pkg/internal/databases/maindb"

	"github.com/gitpushy/pave/cmd/pave/templates/pkg/web/graph"
	"github.com/gitpushy/pave/cmd/pave/templates/pkg/web/graph/generated"
	"github.com/gitpushy/pave/cmd/pave/templates/pkg/web/graph/implementation"
	"github.com/rs/cors"
	"go.uber.org/zap"
)

// RunServer runs the web server
func RunServer(
	ctx context.Context,
	config *configurations.Config,
	parentLogger *logging.Logger,
) error {
	logger := parentLogger.Named("web")
	webConfig := config.Web
	dbConfig := config.Databases.MainDB

	mainDB, err := maindb.Open(ctx, logger, dbConfig)
	if err != nil {
		return err
	}

	resolver := &graph.Resolver{
		Resolver: implementation.Resolver{
			MainDB: mainDB,
		},
	}

	gsrv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: resolver,
	}))

	gsrv.AroundOperations(func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		op := graphql.GetOperationContext(ctx)
		logger := logger.Named(op.OperationName)
		logger.Info("handling")
		defer logger.Info("handled")
		return next(implementation.WithTokenContextFromHeader(ctx, op.Headers))
	})

	c := cors.New(cors.Options{
		AllowedOrigins: webConfig.AllowedOrigins,
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true, // allow Authorization header
	})
	handleCORS := c.Handler

	routes := http.NewServeMux()
	routes.Handle("/graphql", gsrv)
	routes.HandleFunc("/", homeRoute)

	rootHandler := handleCORS(routes)
	port := webConfig.Port
	logger.Info("running web server", zap.Int("port", port))
	return http.ListenAndServe(fmt.Sprintf(":%d", port), rootHandler)
}

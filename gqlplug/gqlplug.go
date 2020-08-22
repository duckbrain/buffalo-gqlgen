package gqlplug

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gobuffalo/buffalo"
)

type schemaResource struct {
	buffalo.BaseResource
	schemaHandler http.Handler
}

func (r schemaResource) List(c buffalo.Context) error {
	h := playground.Handler("GraphQL playground", c.Request().URL.Path)
	h.ServeHTTP(c.Response(), c.Request())
	return nil
}

func (r schemaResource) Create(c buffalo.Context) error {
	r.schemaHandler.ServeHTTP(c.Response(), c.Request())
	return nil
}

func Handler(s graphql.ExecutableSchema) buffalo.Resource {
	sh := handler.NewDefaultServer(s)
	return schemaResource{schemaHandler: sh}
}

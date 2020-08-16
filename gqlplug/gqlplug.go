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
	playgroundHandler http.Handler
	schemaHandler     http.Handler
}

func (r schemaResource) List(c buffalo.Context) error {
	r.playgroundHandler.Handle(c.Response(), c.Request())
	return nil
}

func (r schemaResource) Create(c buffalo.Context) error {
	r.schemaHandler.Handle(c.Response(), c.Request())
	return nil
}

func Handler(s graphql.ExecutableSchema) buffalo.Resource {
	ph := playground.Handler("GraphQL playground", "/")
	sh := handler.NewDefaultServer(s)
	return schemaResource{playgroundHandler: ph, schemaHandler: sh}
}

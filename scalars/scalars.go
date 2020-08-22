package scalars

import (
	"io"
	"strconv"
	"strings"

	"github.com/99designs/gqlgen/graphql"
)

var All = map[string][]string{
	"ID":      {"#UUID", "#NullUUID", "#NullsUUID"},
	"Time":    {"#Time", "#NullsTime"},
	"Date":    {"#Time", "#NullsTime"},
	"Weekday": {"#Weekday"},
	"String":  {"github.com/99designs/gqlgen/graphql.String", "#NullsString"},
	"Int": {
		"github.com/99designs/gqlgen/graphql.Int", "github.com/99designs/gqlgen/graphql.Uint32",
		"github.com/99designs/gqlgen/graphql.Int32", "github.com/99designs/gqlgen/graphql.Int64",
		"#NullsInt", "#NullsUint32", "#NullsInt32", "#NullsInt64",
	},
	"Float": {
		"github.com/99designs/gqlgen/graphql.Float32", "github.com/99designs/gqlgen/graphql.Float64",
		"#NullsFloat32", "#NullsFloat64",
	},
	"Boolean": {"github.com/99designs/gqlgen/graphql.Boolean", "#NullsBool"},
}

func init() {
	// Add prefixes
	for _, s := range All {
		for i := range s {
			if strings.HasPrefix(s[i], "#") {
				s[i] = "github.com/duckbrain/buffalo-gqlgen/scalars." + s[i][1:]
			}
		}
	}
}

func outputWriter(s string) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, err := w.Write([]byte(s))
		if err != nil {
			panic(err)
		}
	})
}

func stringWriter(s string) graphql.Marshaler {
	return outputWriter(strconv.Quote(s))
}

var nullWriter = outputWriter("null")

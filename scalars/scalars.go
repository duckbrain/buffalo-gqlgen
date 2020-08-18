package scalars

import (
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
)

var All = map[string][]string{
	"ID":      {"UUID", "NullUUID", "NullsUUID"},
	"Time":    {"Time", "NullsTime"},
	"Date":    {"Time", "NullsTime"},
	"Weekday": {"Weekday"},
}

func init() {
	// Add prefixes
	for _, s := range All {
		for i := range s {
			s[i] = "github.com/duckbrain/buffalo-gqlgen/scalars." + s[i]
		}
	}
}

func stringWriter(s string) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, err := w.Write([]byte(strconv.Quote(s)))
		if err != nil {
			panic(err)
		}
	})
}

var nullWriter graphql.Marshaler = graphql.WriterFunc(func(w io.Writer) {
	_, err := w.Write([]byte("null"))
	if err != nil {
		panic(err)
	}
})

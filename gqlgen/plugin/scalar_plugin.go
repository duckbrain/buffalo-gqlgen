package plugin

import (
	"log"

	"github.com/99designs/gqlgen/codegen/config"
	"github.com/duckbrain/buffalo-gqlgen/scalars"
	"github.com/vektah/gqlparser/v2/ast"
)

// Scalar impelents a gqlgen plugin that adds implementations for common scalars
// if the source doesn't define them.
type Scalar struct{}

func (Scalar) Name() string {
	return "Scalar"
}
func (Scalar) MutateConfig(cfg *config.Config) error {
	for n, s := range scalars.All {
		t, ok := cfg.Schema.Types[n]
		if ok && t.Kind == ast.Scalar {
			for i := range s {
				cfg.Models.Add(n, s[i])
			}
		}
	}
	return nil
}

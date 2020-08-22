package plugin

import (
	"github.com/99designs/gqlgen/codegen"
	"github.com/99designs/gqlgen/codegen/config"
	"github.com/99designs/gqlgen/plugin/modelgen"
)

type PopGen struct {
}

func (p PopGen) MutateConfig(cfg *config.Config) error {
	return nil
	// for n, t := range cfg.Schema.Types {
	// 	d := t.Directives.ForName("pop")
	// 	if d == nil {
	// 		continue
	// 	}
	// 	modelName := n
	// 	modelArg := d.Arguments.ForName("model")
	// 	if modelArg != nil {
	// 		modelName = modelArg.Value.Raw
	// 	}
	// 	model := cfg.Models[modelName]
	// 	if model == nil {
	// 		return nil
	// 	}
	// 	cfg.
	// }
}
func (p PopGen) GenerateCode(d *codegen.Data) error {
	return nil
	// f, err := os.Create(filepath.Join(d.Config.Resolver.Dir(), "buffalo_pop.resolvers.go"))
	// if err != nil {
	// 	return err
	// }
	// defer f.Close()
	// _, err = f.WriteString("package " + d.Config.Resolver.Package)
	// if err != nil {
	// 	return err
	// }
	// for _, o := range d.Objects {
	// 	o.Fields[0].IsResolver
	// }
	// return err
}

// Defining mutation function
func mutatePopModelHook(b *modelgen.ModelBuild) *modelgen.ModelBuild {
	for _, model := range b.Models {
		for _, field := range model.Fields {
			field.Tag += ` db:"` + model.Name + `.` + field.Name + `"`
		}
	}

	return b
}

func NewPopModelGen() *modelgen.Plugin {
	return &modelgen.Plugin{
		MutateHook: mutatePopModelHook,
	}
}

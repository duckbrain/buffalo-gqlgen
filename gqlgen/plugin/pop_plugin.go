package plugin

import (
	"syscall"

	"github.com/99designs/gqlgen/codegen/config"
	"github.com/99designs/gqlgen/codegen/templates"
)

type PopGen struct {
	Filename string
}

func (p PopGen) Name() string {
	return "Pop"
}

func (p PopGen) MutateConfig(cfg *config.Config) error {
	_ = syscall.Unlink(p.Filename)

	return templates.Render(templates.Options{
		PackageName: "models",
		Filename:    p.Filename,
		Data: &PopBuild{
			Config: cfg,
		},
		GeneratedHeader: true,
		Packages:        cfg.Packages,
	})
}

type PopBuild struct {
	*config.Config
}

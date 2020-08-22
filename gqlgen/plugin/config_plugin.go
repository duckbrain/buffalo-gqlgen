package plugin

import (
	"errors"

	"github.com/99designs/gqlgen/codegen/config"
	"golang.org/x/tools/go/packages"
)

type Config struct{}

func (Config) Name() string {
	return "Config"
}

func (Config) MutateConfig(cfg *config.Config) error {
	pkgs, err := packages.Load(&packages.Config{
		Dir:  "./models",
		Mode: packages.NeedName,
	})
	if err != nil {
		return err
	}
	if len(pkgs) == 0 {
		return errors.New("No packages found in ./models")
	}
	pkgPath := ""
	for _, pkg := range pkgs {
		if pkg.Name == "models" {
			pkgPath = pkg.PkgPath
			break
		}
	}

	cfg.Model.Filename = "models.models_gen.go"
	cfg.Model.Package = "models"
	cfg.Resolver.Layout = config.LayoutFollowSchema
	cfg.Resolver.DirName = "actions/gql"
	cfg.Resolver.Package = "gql"
	cfg.OmitSliceElementPointers = true
	cfg.AutoBind = append(cfg.AutoBind, pkgPath)
	return nil
}

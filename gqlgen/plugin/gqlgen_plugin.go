package plugin

import (
	"github.com/99designs/gqlgen/codegen"
	"github.com/99designs/gqlgen/codegen/config"
	"github.com/99designs/gqlgen/plugin"
	"github.com/pkg/errors"
)

// TODO plugin that auto implements resolvers using pop

type Plugins []plugin.Plugin

func (Plugins) Name() string {
	return "Buffalo"
}

func New() *Plugins {
	return &Plugins{
		Config{},
		Scalar{},
		NewPopModelGen(),
	}
}

func (p Plugins) MutateConfig(cfg *config.Config) error {
	for i := range p {
		if x, ok := p[i].(plugin.ConfigMutator); ok {
			err := x.MutateConfig(cfg)
			if err != nil {
				return errors.Errorf("%v: %v", p[i].Name(), err)
			}
		}
	}
	return nil
}

func (p Plugins) GenerateCode(cfg *codegen.Data) error {
	for i := range p {
		if x, ok := p[i].(plugin.CodeGenerator); ok {
			err := x.GenerateCode(cfg)
			if err != nil {
				return errors.Errorf("%v: %v", p[i].Name(), err)
			}
		}
	}
	return nil
}

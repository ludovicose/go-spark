package config

import (
	"github.com/ludovicose/go-spark/pkg/facades"
)

type ServiceProvider struct {
}

func (config *ServiceProvider) Boot() {
}

func (config *ServiceProvider) Register() {
	app := Application{}
	facades.Config = app.Init()
}
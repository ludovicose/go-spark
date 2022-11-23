package http

import "github.com/ludovicose/go-spark/pkg/facades"

type ServiceProvider struct {
}

func (database *ServiceProvider) Boot() {
}

func (database *ServiceProvider) Register() {
	app := Application{}
	facades.Request, facades.Response = app.Init()
}
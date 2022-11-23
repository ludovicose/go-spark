package route

import (
	"github.com/ludovicose/go-spark/pkg/facades"
)

type ServiceProvider struct {
}

func (route *ServiceProvider) Boot() {

}

func (route *ServiceProvider) Register() {
	app := Application{}
	facades.Route = app.Init()
}

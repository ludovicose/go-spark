package providers

import (
	"github.com/ludovicose/go-spark/internal/http"
	"github.com/ludovicose/go-spark/internal/routes"
	"github.com/ludovicose/go-spark/pkg/facades"
)

type RouteServiceProvider struct {
}

func (receiver *RouteServiceProvider) Boot() {
	//Add HTTP middlewares.
	facades.Route.Use(http.Kernel{}.Middleware()...)

	//Add routes
	routes.Web()
}

func (receiver *RouteServiceProvider) Register() {

}

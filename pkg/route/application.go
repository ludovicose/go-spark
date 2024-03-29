package route

import (
	"github.com/gin-gonic/gin"
	"github.com/ludovicose/go-spark/pkg/facades"
)

type Application struct {
}

func (app *Application) Init() *gin.Engine {
	//When running in console, don't show gin log.
	if facades.Config.GetBool("app.debug") {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	route := gin.New()

	return route
}

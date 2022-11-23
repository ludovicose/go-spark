package core

import (
	"github.com/ludovicose/go-spark/pkg/foundation"
	"github.com/ludovicose/go-spark/config"
)

func Boot(){
	app := foundation.Application{}

	app.Boot()

	config.Boot()
}

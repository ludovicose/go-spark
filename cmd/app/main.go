package main

import (
	"github.com/ludovicose/go-spark/core"
	"github.com/ludovicose/go-spark/pkg/facades"
)

func main() {
	core.Boot()

	go facades.Route.Run(facades.Config.GetString("app.host"))

	select {}
}

package config

import (
	"github.com/ludovicose/go-spark/internal/providers"
	"github.com/ludovicose/go-spark/pkg/contracts"
	"github.com/ludovicose/go-spark/pkg/facades"
	"github.com/ludovicose/go-spark/pkg/http"
	"github.com/ludovicose/go-spark/pkg/route"
)

func Boot() {}

func init() {
	config := facades.Config
	config.Add("app", map[string]interface{}{
		"name": config.Env("APP_NAME", "nft"),

		"env": config.Env("APP_ENV", "production"),

		"debug": config.Env("APP_DEBUG", false),

		"key": config.Env("APP_KEY", ""),

		"url": config.Env("APP_URL", "http://localhost"),

		"host": config.Env("APP_HOST", "127.0.0.1:3000"),

		"grpc_host": config.Env("GRPC_HOST", ""),

		"providers": []contracts.ServiceProvider{
			&http.ServiceProvider{},
			&route.ServiceProvider{},
			&providers.RouteServiceProvider{},
		},
	})
}

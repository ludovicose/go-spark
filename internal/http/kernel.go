package http

import (
	"github.com/gin-gonic/gin"
	"github.com/ludovicose/go-spark/internal/http/middleware"
)

type Kernel struct {
}

func (kernel Kernel) Middleware() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		gin.Logger(),
		middleware.Cors(),
	}
}

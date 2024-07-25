package main

import (
	"github.com/freightcms/api-template/api"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	v1Routes := server.Group("/v1")

	api.Register(v1Routes)

	if err := server.Run(":3001"); err != nil {
		panic(err)
	}
}

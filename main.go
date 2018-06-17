package main

import (
	"github.com/dashotv/server/actions"
	"github.com/gin-gonic/gin"
)

const JWT_SECRET = "blarg blarg blarg"

func main() {
	router := gin.Default()
	router.GET("/", actions.Home)

	api := router.Group("/api")
	actions.ReleaseRoutes(api)

	router.Run(":8080")
}

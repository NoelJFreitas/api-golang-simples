package main

import (
	"github.com/NoelJFreitas/api-golang/api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	routes.AppRoutes(app)

	app.Run("localhost:8000")
}

package main

import (
	"github.com/Shakti-Tamang/crud-app/initializers"
	"github.com/Shakti-Tamang/crud-app/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariable()
	initializers.ConnectDB()

}
func main() {

	r := gin.Default()
	routes.ItemsRoutes(r)

	r.Run()
}

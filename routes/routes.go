package routes

import (
	"github.com/Shakti-Tamang/crud-app/controllers"
	"github.com/gin-gonic/gin"
)

func ItemsRoutes(r *gin.Engine) {
	useGroup := r.Group("/create")

	{

		useGroup.POST("", controllers.CreateItem)
		useGroup.GET("", controllers.GetAll)

	}

}

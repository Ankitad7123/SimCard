package routes

import (
	"simcard/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UrlPath(r *gin.Engine, db *gorm.DB) {
	views := controllers.NewDB{DB: db}

	r.POST("/activate", views.CreateSim)           // activate api
	r.POST("/deactivate", views.Deactivate)        //deactivate api
	r.GET("/simdetails/:simnumber", views.Handler) // get the details

}

package routes

import (
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

type ApiV1Handler interface {
	GetAll(ctx *gin.Context)
	GetById(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

func init() {
	router = gin.Default()
}

func SetBasicCRUD(group string, collectionHandler ApiV1Handler) {
	gp := router.Group(group)
	gp.GET("/listing", collectionHandler.GetAll)
	gp.GET("/listing/:id", collectionHandler.GetById)
	gp.POST("/listing", collectionHandler.Create)
	gp.PUT("/listing/:id", collectionHandler.Update)
	gp.DELETE("/listing/:id", collectionHandler.Delete)
}

func Serve(addr ...string) {
	router.Run(addr...)
}

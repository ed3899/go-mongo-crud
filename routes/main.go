package routes

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	router = gin.Default()
}

func Serve(addr ...string) {
	router.Run(addr...)
}

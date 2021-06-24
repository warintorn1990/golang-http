package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetUp(router *gin.Engine) {
	authenAPI := router.Group("/authen")
	{
		authenAPI.GET("/login", Login)
		authenAPI.GET("/register", Register)
		authenAPI.GET("/profile", func(c *gin.Context) {
			c.String(http.StatusOK, "Profile")
		})
	}

	stockAPI := router.Group("/stock")
	{
		stockAPI.GET("/list", ListProduct)
		stockAPI.GET("/create", CreateProduct)
	}

}

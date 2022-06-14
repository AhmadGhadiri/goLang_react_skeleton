package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setRouter() *gin.Engine {
	// Creates default gin router with Logger and Recovery middleware already attached
	router := gin.Default()

	// Create API route group
	api := router.Group("/api")
	api.POST("/signup", signUp)
	api.POST("/signin", signIn)

	// Create a group for authorized users
	authorized := api.Group("/")
	authorized.Use(authorization)
	{
		authorized.POST("/posts", createPost)
		authorized.GET("/posts", indexPosts)
		authorized.PUT("/posts", updatePost)
		authorized.DELETE("/posts/:id", deletePost)
	}
	// {
	// 	// Add /hello GET route to router and define route handler function
	// 	api.GET("/hello", func(ctx *gin.Context) {
	// 		ctx.JSON(200, gin.H{"msg": "world"})
	// 	})
	// }

	router.NoRoute(func(ctx *gin.Context) { ctx.JSON(http.StatusNotFound, gin.H{}) })

	return router
}

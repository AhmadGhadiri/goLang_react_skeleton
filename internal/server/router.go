package server

import (
	"net/http"
	"rgb/internal/store"

	"github.com/gin-gonic/gin"
)

func setRouter() *gin.Engine {
	// Creates default gin router with Logger and Recovery middleware already attached
	router := gin.Default()

	router.Use(CORSMiddleware())

	// Create API route group
	api := router.Group("/api")
	api.Use(customErrors)
	{
		/* bind request data before even hitting signUp and signIn handlers,
		   which means that handlers will only be reached if form validations
		   are passed. With setup like this, handlers don’t need to think
		   about binding errors, because there was none if handler is reached
		*/
		api.POST("/signup", gin.Bind(store.User{}), signUp)
		api.POST("/signin", gin.Bind(store.User{}), signIn)

		// Test functions
		api.GET("/test/all", testAll)
		api.GET("/test/user", testUser)
	}

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

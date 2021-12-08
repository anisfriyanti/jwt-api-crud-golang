package main

import (
	"jwtapi/controllers"
	"jwtapi/models"
	"jwtapi/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Connect to database
	models.ConnectDatabase()

	// Routes
		r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})


	api := r.Group("/api")
	{
		public := api.Group("/public")
		{
			public.POST("/login", controllers.Login)
			public.POST("/signup", controllers.Signup)
			public.GET("/test",controllers.Calltes)


			products := api.Group("/data").Use(middlewares.Authz())
		{
			products.POST("/transaction", controllers.CreateTransaction)
			products.POST("/pay", controllers.UpdateTransaction)
			products.GET("/product", controllers.FindProducts)
			products.POST("/product", controllers.CreateProduct)
			products.PATCH("/product/:id", controllers.UpdateProduct)
			products.DELETE("/product/:id", controllers.DeleteProduct)


		}

			
			
		}

		
		
	}
	
	// Run the server
	r.Run()
}

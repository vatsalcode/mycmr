// @author- Vatsal Verma
package main

import (
	"mycrm/controllers"
	"mycrm/services"
	"mycrm/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	db := utils.ConnectDB()
	jwtSecret := "your_jwt_secret" // Replace with a secure secret

	userService := services.NewUserService(db)
	customerService := services.NewCustomerService(db)
	interactionService := services.NewInteractionService(db)
	authService := services.NewAuthService(db, jwtSecret)

	userController := controllers.NewUserController(userService)
	customerController := controllers.NewCustomerController(customerService)
	interactionController := controllers.NewInteractionController(interactionService)
	authController := controllers.NewAuthController(authService)

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to the CRM System by Vatsal"})
	})

	// Auth routes
	router.POST("/register", authController.Register)
	router.POST("/login", authController.Login)

	// User routes
	router.POST("/users", userController.CreateUser)
	router.GET("/users/:id", userController.GetUser)
	router.PUT("/users/:id", userController.UpdateUser)
	router.DELETE("/users/:id", userController.DeleteUser)

	// Customer routes
	router.POST("/customers", customerController.CreateCustomer)
	router.GET("/customers/:id", customerController.GetCustomer)
	router.PUT("/customers/:id", customerController.UpdateCustomer)
	router.DELETE("/customers/:id", customerController.DeleteCustomer)

	// Interaction routes
	router.POST("/interactions", interactionController.CreateInteraction)
	router.GET("/customers/:id/interactions", interactionController.GetInteractions)
	router.GET("/interactions/stats", interactionController.GetInteractionStats)

	router.Run(":8080")
}

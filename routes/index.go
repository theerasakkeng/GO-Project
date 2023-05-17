package routes

import (
	"project-study/db"
	"project-study/handler"
	"project-study/repository"
	"project-study/service"

	"github.com/gin-gonic/gin"
)

var Router = gin.Default()

func InitRoute() {
	db.InitConfig()
	db, _ := db.InitDB()

	customerRepository := repository.NewCustomerRepositoryDB(db)
	customerService := service.NewCustomerService(customerRepository)
	customerHandler := handler.NewCustomerHandler(customerService)

	customerRoutes := Router.Group("/api/customer")
	{
		customerRoutes.GET("/getcustomers", customerHandler.GetCustomerList)
		customerRoutes.GET("/getcustomerdetail", customerHandler.GetCustomerDetail)
	}
}

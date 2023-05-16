package main

import (
	"project-study/handler"
	"project-study/repository"
	"project-study/service"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func main() {
	dsn := "sqlserver://sa:Keng1234@localhost?database=TestDB"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	customerRepository := repository.NewCustomerRepositoryDB(db)
	customerService := service.NewCustomerService(customerRepository)
	customerHandler := handler.NewCustomerHandler(customerService)

	r := gin.Default()

	customerRoutes := r.Group("/api/customer")
	{
		customerRoutes.GET("/getcustomers", customerHandler.GetCustomerList)
		customerRoutes.GET("/getcustomerdetail", customerHandler.GetCustomerDetail)
	}

	r.Run()
}

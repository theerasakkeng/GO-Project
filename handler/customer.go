package handler

import (
	"net/http"
	"project-study/repository"
	"project-study/service"

	"github.com/gin-gonic/gin"
)

type customerHandler struct {
	custSrv service.CustomerService
}

func NewCustomerHandler(custSrv service.CustomerService) customerHandler {
	return customerHandler{custSrv: custSrv}

}

func (h customerHandler) GetCustomerList(c *gin.Context) {
	customers, err := h.custSrv.GetCustomerList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": customers,
	})
}

func (h customerHandler) GetCustomerDetail(c *gin.Context) {
	id := c.Query("id")
	customer, err := h.custSrv.GetCustomerDetail(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": customer,
	})
}

func (h customerHandler) InsertCustomer(c *gin.Context) {
	var customer repository.CustomerRequest
	err := c.BindJSON(&customer)
	if err != nil {
		panic(err)
	}
	res, err := h.custSrv.InsertCustomer(customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}

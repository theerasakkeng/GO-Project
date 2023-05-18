package service

import "project-study/repository"

type CustomerResponse struct {
	Customer_Id int     `json:"customer_id"`
	First_Name  string  `json:"first_name"`
	Last_Name   string  `json:"last_name"`
	Phone       *string `json:"phone"`
}

type CustomerService interface {
	GetCustomerList() ([]CustomerResponse, error)
	GetCustomerDetail(string) (*CustomerResponse, error)
	InsertCustomer(repository.CustomerRequest) (*repository.CustomerRequest, error)
}

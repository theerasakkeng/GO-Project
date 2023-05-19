package repository

import "gorm.io/gorm"

type customerRepositoryDB struct {
	db *gorm.DB
}

func NewCustomerRepositoryDB(db *gorm.DB) customerRepositoryDB {
	return customerRepositoryDB{db: db}

}

func (r customerRepositoryDB) GetAll() ([]Customer, error) {
	var customers []Customer
	result := r.db.Find(&customers)
	if result.Error != nil {
		return nil, result.Error
	}
	return customers, nil
}

func (r customerRepositoryDB) GetById(id string) (*Customer, error) {
	var customer Customer
	result := r.db.Raw("select * from sales.customers where customer_id = ?", id).Scan(&customer)
	if result.Error != nil {
		return nil, result.Error
	}
	return &customer, nil
}

func (r customerRepositoryDB) PostCustomer(customer CustomerRequest) (*CustomerRequest, error) {
	var customerReq = customer
	result := r.db.Create(&customerReq)
	if result.Error != nil {
		return nil, result.Error
	}
	return &customerReq, nil
}

func (r customerRepositoryDB) UpdateCustomer(customer CustomerRequest, id string) (*CustomerRequest, error) {
	var customerReq = customer
	result := r.db.Where("customer_id = ?", id).Save(&customerReq)
	if result.Error != nil {
		return nil, result.Error
	}
	return &customerReq, nil
}

func (r customerRepositoryDB) DeleteCustomerById(id int) (string, error) {
	var customerReq = Customer{}
	result := r.db.Where("customer_id = ?", id).Delete(&customerReq)
	if result.Error != nil {
		return "error", result.Error
	}
	return "success", nil
}

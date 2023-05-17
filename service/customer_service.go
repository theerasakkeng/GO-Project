package service

import "project-study/repository"

type customerService struct {
	custRepo repository.CustomerRepository
}

func NewCustomerService(custRepo repository.CustomerRepository) customerService {
	return customerService{custRepo: custRepo}
}

func (s customerService) GetCustomerList() ([]CustomerResponse, error) {
	customers, err := s.custRepo.GetAll()
	if err != nil {
		return nil, err
	}
	custResponses := []CustomerResponse{}
	for _, customer := range customers {
		custResponse := CustomerResponse{
			Customer_Id: customer.Customer_Id,
			First_Name:  customer.First_Name,
			Last_Name:   customer.Last_Name,
			Phone:       customer.Phone,
		}
		custResponses = append(custResponses, custResponse)
	}
	return custResponses, nil
}

func (s customerService) GetCustomerDetail(id string) (*CustomerResponse, error) {
	customer, err := s.custRepo.GetById(id)
	if err != nil {
		return nil, err
	}
	custResponse := CustomerResponse{
		Customer_Id: customer.Customer_Id,
		First_Name:  customer.First_Name,
		Last_Name:   customer.Last_Name,
		Phone:       customer.Phone,
	}
	return &custResponse, nil
}

func (s customerService) InsertCustomer(customerreq repository.CustomerRequest) (string, error) {
	_, err := s.custRepo.PostCustomer(customerreq)
	if err != nil {
		return "error", err
	}
	return "success", nil
}

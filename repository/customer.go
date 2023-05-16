package repository

type Customer struct {
	Customer_Id int     `db:"customer_id"`
	First_Name  string  `db:"first_name"`
	Last_Name   string  `db:"last_name"`
	Phone       *string `db:"phone"`
	Email       string  `db:"email"`
	Street      string  `db:"street"`
	City        string  `db:"city"`
	State       string  `db:"state"`
	Zip_Code    string  `db:"zip_code"`
}

type CustomerRepository interface {
	GetAll() ([]Customer, error)
	GetById(string) (*Customer, error)
}

type Tabler interface {
	TableName() string
}

func (Customer) TableName() string {
	return "sales.customers"
}

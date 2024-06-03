package model

type Users struct {
	Id       int
	Username string
	Email    string
	Password string
}

type Products struct{
	Id int
	Name string
	Description string
	Price float64
	Stock_quantity int
}
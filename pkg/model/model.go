package model

type User struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"email"`
}
type Customer struct {
	Name      string `json:"name" validate:"required"`
	Cuit      string `json:"name" validate:"required"`
	Direccion string `json:"name" validate:"required"`
	Telefono  string `json:"name" validate:"phone"`
}
type Resp struct {
	Status  string `json:status`
	Mensaje string `json:mensaje`
}

type RelUSerCustome struct {
	User
	Customer
}

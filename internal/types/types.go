package types

type Student struct{
	Id int `json:"id"`
	Name string	`json:"name" validate:"required"`
	Email string	`json:"email" validate:"email"`
	Age  int		`json:"age" validate:"required"`
}
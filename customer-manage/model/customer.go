package model

import "fmt"

type Customer struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Email  string `json:"email"`
}

func NewCustomer(id int, name, gender, email string) *Customer {
	return &Customer{
		Id:     id,
		Name:   name,
		Gender: gender,
		Email:  email,
	}
}

func CustomerFormat(c *Customer) string {
	return fmt.Sprintf("\t%v\t\t%v\t\t%v\t\t%v\n", c.Id, c.Name, c.Gender, c.Email)
}

package model

import "fmt"

// Customer model 结构体
type Customer struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Email  string `json:"email"`
}

// NewCustomer 工厂模式，New 一个 Customer
func NewCustomer(id int, name, gender, email string) *Customer {
	return &Customer{
		ID:     id,
		Name:   name,
		Gender: gender,
		Email:  email,
	}
}

// CustomerFormat 格式化输出
func CustomerFormat(c *Customer) string {
	return fmt.Sprintf("\t%v\t\t%v\t\t%v\t\t%v\n", c.ID, c.Name, c.Gender, c.Email)
}

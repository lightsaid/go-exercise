package service

import (
	"fmt"
	"go-exercise/customer-manage/model"
)

type CustomerService struct {
	ID           int
	CustomerList []model.Customer
}

func (c *CustomerService) GetCustomer() string {
	var str string
	for _, v := range c.CustomerList {
		str += model.CustomerFormat(&v)
	}
	return str
}

func (c *CustomerService) Update(id int, name, gender, email string) bool {
	index, customer := c.GetCustomerById(id)
	if customer != nil && customer.Id > 0 {
		customer.Name = name
		customer.Gender = gender
		customer.Email = email
		c.CustomerList[index] = *customer
		return true
	} else {
		fmt.Println("没有查找到Customer")
		return false
	}
}

func (c *CustomerService) Delete(id int) bool {
	index, customer := c.GetCustomerById(id)
	if customer != nil && customer.Id > 0 {
		c.CustomerList = append(c.CustomerList[:index], c.CustomerList[index+1:]...)
		return true
	} else {
		fmt.Println("没有查找到Customer")
		return false
	}
}

func (c *CustomerService) GetCustomerById(id int) (index int, customer *model.Customer) {
	for k, v := range c.CustomerList {
		if v.Id == id {
			index = k
			customer = &v
			break
		}
	}
	return
}

func (c *CustomerService) AddCustomer(customer *model.Customer) {
	c.ID++
	customer.Id = c.ID
	c.CustomerList = append(c.CustomerList, *customer)
	fmt.Println("List:", c.CustomerList)
}

func (this *CustomerService) GetCustomersList() (customers []model.Customer) {
	return this.CustomerList
}

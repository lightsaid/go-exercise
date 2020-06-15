package service

import (
	"fmt"
	"go-exercise/customer-manage/dao"
	"go-exercise/customer-manage/model"
)

// CustomerService server 层结构体
type CustomerService struct {
	ID           int
	CustomerList []model.Customer
	CustomerDao  *dao.CustomerController
}

// Add 新增
func (service *CustomerService) Add() error {
	customer := &model.Customer{}
	fmt.Println("请输入Name:")
	fmt.Scanln(&customer.Name)
	fmt.Println("请输入Gender:")
	fmt.Scanln(&customer.Gender)
	fmt.Println("请输入Email:")
	fmt.Scanln(&customer.Email)
	return nil
}

// List 获取列表
func (service *CustomerService) List() string {
	var str string
	conn := service.CustomerDao.Pool.Get()
	defer conn.Close()
	conn.Do("HGet", "customers")
	for _, v := range service.CustomerList {
		str += model.CustomerFormat(&v)
	}
	return str
}

// Update 更新数据
func (service *CustomerService) Update() bool {
	var id int
	var (
		name, gender, email string
	)
	index, customer := service.GetCustomerByID(id)
	if customer != nil && customer.Id > 0 {
		customer.Name = name
		customer.Gender = gender
		customer.Email = email
		service.CustomerList[index] = *customer
		return true
	} else {
		fmt.Println("没有查找到Customer")
		return false
	}
}

// Delete 删除
func (service *CustomerService) Delete() bool {
	var id int
	index, customer := service.GetCustomerByID(id)
	if customer != nil && customer.Id > 0 {
		service.CustomerList = append(service.CustomerList[:index], service.CustomerList[index+1:]...)
		return true
	} else {
		fmt.Println("没有查找到Customer")
		return false
	}
}

// GetCustomerByID 获取一个customer
func (service *CustomerService) GetCustomerByID(id int) (index int, customer *model.Customer) {
	for k, v := range service.CustomerList {
		if v.Id == id {
			index = k
			customer = &v
			break
		}
	}
	return
}

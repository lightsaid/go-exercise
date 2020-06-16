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
func (service *CustomerService) Add() {
	customer := &model.Customer{}
	fmt.Println("请输入Name:")
	fmt.Scanln(&customer.Name)
	fmt.Println("请输入Gender:")
	fmt.Scanln(&customer.Gender)
	fmt.Println("请输入Email:")
	fmt.Scanln(&customer.Email)

	err := service.CustomerDao.AddCustomer(customer)
	if err != nil {
		fmt.Println("Add Failed:", err)
	}
	fmt.Println("Add Success")
}

// List 获取列表
func (service *CustomerService) List() string {
	var str string
	service.CustomerDao.GetCustomerList()
	// for _, v := range service.CustomerList {
	// 	str += model.CustomerFormat(&v)
	// }
	return str
}

// Update 更新数据
func (service *CustomerService) Update() {
	var id int
	var (
		name, gender, email string
	)
	fmt.Println("Id:")
	fmt.Scanf("%d\n", &id)
	fmt.Println("Name:")
	fmt.Scanln(&name)
	fmt.Println("Gender:")
	fmt.Scanln(&gender)
	fmt.Println("Email:")
	fmt.Scanln(&email)
	// 先查找是否存在
	customer, err := service.GetCustomerByID(id)
	if err != nil {
		fmt.Println("not find customer", err)
		return
	}
	if customer != nil && customer.Id > 0 {
		customer.Name = name
		customer.Gender = gender
		customer.Email = email
		err := service.CustomerDao.UpdateCustomer(customer)
		if err != nil {
			fmt.Println("update fail:", err)
			return
		}
		fmt.Println("update success!")
	} else {
		fmt.Println("not find customer")
		return
	}
}

// Delete 删除
func (service *CustomerService) Delete() {
	var id int
	fmt.Println("请输入Id:")
	fmt.Scanf("%d\n", &id)
	err := service.CustomerDao.DeleteCustomerById(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Delete Success!")
}

// GetCustomerByID 获取一个customer
func (service *CustomerService) GetCustomerByID(id int) (customer *model.Customer, err error) {
	customer, err = service.CustomerDao.GetCustomerByID(id)
	if err != nil {
		return
	}
	return
}

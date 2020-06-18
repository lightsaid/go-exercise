package service

import (
	"fmt"
	"go-exercise/customer-manage/dao"
	"go-exercise/customer-manage/model"
	"os"
	"strings"
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
func (service *CustomerService) List() {
	var str string
	customers, err := service.CustomerDao.GetCustomerList()
	if err != nil {
		fmt.Println("GetCustomerList fail:", err)
		return
	}
	for _, v := range customers {
		str += model.CustomerFormat(&v)
	}
	fmt.Println(str)
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
	if customer != nil && customer.ID > 0 {
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
	err := service.CustomerDao.DeleteCustomerByID(id)
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

// Exit 确认退出
func (service *CustomerService) Exit() {
	var key string
	fmt.Println("确定要退出？（y/n）")
	fmt.Scanf("%s\n", &key)
	if strings.ToUpper(key) == "Y" {
		fmt.Println("正在退出系统。。。")
		os.Exit(0)
	}
}

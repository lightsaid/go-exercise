package dao

import (
	"encoding/json"
	"errors"
	"fmt"

	"go-exercise/customer-manage/model"

	"github.com/gomodule/redigo/redis"
)

// CustomerController 控制器结构体
type CustomerController struct {
	Pool *redis.Pool
}

// AddCustomer 添加一个customer
func (controller *CustomerController) AddCustomer(customer *model.Customer) (err error) {
	// 取出一根连接池
	conn := controller.Pool.Get()
	defer conn.Close()
	// 设置customerId 自增 1
	id, err := redis.Int(conn.Do("HINCRBY", "customers", "customerId", 1))
	if err != nil {
		fmt.Println("Hash HINCRBY err:", err)
		return
	}
	fmt.Println("id:", id)

	customer.ID = id
	str, err := json.Marshal(customer)

	_, err = conn.Do("HSet", "customers", id, str)
	if err != nil {
		fmt.Println("HSet err:", err)
		return
	}
	return nil
	// conn.Do("HSet", "customers", customer.ID, string(customer))
	// 通过命令设置redis Hash customers 基础数据
	// hset customers 1 "{\"id\":1,\"name\":\"Ly\",\"gender\":\"male\",\"email\":\"Ly@qq.com\"}"
	// hset customers customerId 1 // 设置customerId 为 0
	// HINCRBY customers customerId 1 // 设置customerId 自增 1
}

// GetCustomerByID 获取customer
func (controller *CustomerController) GetCustomerByID(id int) (customer *model.Customer, err error) {
	conn := controller.Pool.Get()
	defer conn.Close()
	result, err := redis.String(conn.Do("HGet", "customers", id))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(result), &customer)
	if err != nil {
		return
	}
	return
}

// DeleteCustomerByID 删除
func (controller *CustomerController) DeleteCustomerByID(id int) error {
	customer, err := controller.GetCustomerByID(id)
	if err != nil {
		return err
	}
	if customer != nil && customer.ID > 0 {
		conn := controller.Pool.Get()
		defer conn.Close()
		_, err = conn.Do("HDel", "customers", id)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("Delete fail, not find customer")
}

// UpdateCustomer 更新
func (controller *CustomerController) UpdateCustomer(customer *model.Customer) (err error) {
	conn := controller.Pool.Get()
	defer conn.Close()
	cid := customer.ID
	str, err := json.Marshal(customer)
	if err != nil {
		return err
	}
	_, err = conn.Do("HSet", "customers", cid, str)
	return err
}

// GetCustomerList 获取列表
func (controller *CustomerController) GetCustomerList() (customers []model.Customer, err error) {
	conn := controller.Pool.Get()
	defer conn.Close()
	res, err := redis.StringMap(conn.Do("HGetAll", "customers"))
	if err != nil {
		fmt.Println(err)
		return
	}
	for k, v := range res {
		customer := model.Customer{}
		if k != "customerId" {
			err = json.Unmarshal([]byte(v), &customer)
			if err != nil {
				continue
			}
			customers = append(customers, customer)
		}
	}
	return
}

package dao

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-exercise/customer-manage/helper"
	"go-exercise/customer-manage/model"
	"strconv"

	"github.com/gomodule/redigo/redis"
)

type CustomerController struct {
	Pool *redis.Pool
}

func (this *CustomerController) AddCustomer(customer *model.Customer) (err error) {
	// 取出一根连接池
	conn := this.Pool.Get()
	defer conn.Close()
	// 设置customerId 自增 1
	id, err := redis.Int(conn.Do("HINCRBY", "customers", "customerId", 1))
	if err != nil {
		fmt.Println("Hash HINCRBY err:", err)
		return
	}
	fmt.Println("id:", id)

	customer.Id = id
	str, err := json.Marshal(customer)

	_, err = conn.Do("HSet", "customers", id, str)
	if err != nil {
		fmt.Println("HSet err:", err)
		return
	}
	return nil
	// conn.Do("HSet", "customers", customer.Id, string(customer))
	// 通过命令设置redis Hash customers 基础数据
	// hset customers 1 "{\"id\":1,\"name\":\"Ly\",\"gender\":\"male\",\"email\":\"Ly@qq.com\"}"
	// hset customers customerId 1 // 设置customerId 为 0
	// HINCRBY customers customerId 1 // 设置customerId 自增 1
}

func (this *CustomerController) GetCustomerByID(id int) (customer *model.Customer, err error) {
	conn := this.Pool.Get()
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

func (this *CustomerController) DeleteCustomerById(id int) error {
	customer, err := this.GetCustomerByID(id)
	if err != nil {
		return err
	}
	if customer != nil && customer.Id > 0 {
		conn := this.Pool.Get()
		defer conn.Close()
		_, err = conn.Do("HDel", "customers", id)
		if err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("Delete Fail Not Find Customer.")
	}
}

func (this *CustomerController) UpdateCustomer(customer *model.Customer) (err error) {
	conn := this.Pool.Get()
	defer conn.Close()
	cid := customer.Id
	str, err := json.Marshal(customer)
	if err != nil {
		return err
	}
	_, err = conn.Do("HSet", "customers", cid, str)
	return err
}

func (this *CustomerController) GetCustomerList() (customers []model.Customer, err error) {
	conn := this.Pool.Get()
	defer conn.Close()
	res, err := redis.StringMap(conn.Do("HGetAll", "customers"))
	if err != nil {
		fmt.Println(err)
		return
	}
	for k, v := range res {
		n, err := strconv.Atoi(k)
		if err != nil {
			continue
		}
		fmt.Println("key:", helper.TypeInference(n))
		if helper.TypeInference(k) == "int" {
			customer := model.Customer{}
			err = json.Unmarshal([]byte(v), &customer)
			fmt.Println("customer=", customer)
			if err != nil {
				return nil, err
			}
			customers = append(customers, customer)
		}
	}
	fmt.Printf("customers=%v", customers)
	return
}

// func (this *CustomerController) GetNextId() int {
// 	conn := this.pool.Get()
// 	defer conn.Close()
// 	return 0
// }

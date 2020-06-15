package dao

import (
	"encoding/json"
	"fmt"
	"go-exercise/customer-manage/model"

	"github.com/gomodule/redigo/redis"
)

type CustomerController struct {
	pool *redis.Pool
}

func (this *CustomerController) AddCustomer(customer *model.Customer) (err error) {
	// 取出一根连接池
	conn := this.pool.Get()
	defer conn.Close()
	// 设置customerId 自增 1
	id, err := redis.Int64(conn.Do("HINCRBY", "customers", "customerId", 1))
	if err != nil {
		fmt.Println("Hash HINCRBY err:", err)
		return
	}
	fmt.Println("id:", id)

	customer.Id = id
	str, err := json.Marshal(customer)

	// conn.Do("HSet", "customers", customer.Id, string(customer))
	// 通过命令设置redis Hash customers 基础数据
	// hset customers 1 "{\"id\":1,\"name\":\"Ly\",\"gender\":\"male\",\"email\":\"Ly@qq.com\"}"
	// hset customers customerId 1 // 设置customerId 为 0
	// HINCRBY customers customerId 1 // 设置customerId 自增 1

	return nil
}

// func (this *CustomerController) GetNextId() int {
// 	conn := this.pool.Get()
// 	defer conn.Close()
// 	return 0
// }

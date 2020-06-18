package main

import (
	"fmt"
	"go-exercise/customer-manage/dao"
	"go-exercise/customer-manage/model"
	"go-exercise/customer-manage/service"
	"time"

	"github.com/gomodule/redigo/redis"
)

// TODO: 1. 客户列表缓存，分页，添加 customer 数据格式校验
// TODO: 2. 返回json api 接口
// application 应用程序结构体
type application struct {
	pool    *redis.Pool
	loop    bool
	key     int
	service *service.CustomerService
}

// 菜单展示
func (app *application) showMenu() {
	for app.loop {
		fmt.Println("-----------------客户信息管理软件-----------------")
		fmt.Println("                 1 添 加 客 户")
		fmt.Println("                 2 修 改 客 户")
		fmt.Println("                 3 删 除 客 户")
		fmt.Println("                 4 客 户 列 表")
		fmt.Println("                 5 退       出")
		fmt.Print("请选择(1-5)：")
		fmt.Scanf("%d\n", &app.key)
		switch app.key {
		case 1:
			app.service.Add()
		case 2:
			app.service.Update()
		case 3:
			app.service.Delete()
		case 4:
			app.service.List()
		case 5:
			app.service.Exit()
		default:
			fmt.Print("输入有误，重新输入：")
		}
	}
}

// 初始化app
func initApp() application {
	service := &service.CustomerService{
		CustomerList: []model.Customer{},
		CustomerDao: &dao.CustomerController{
			Pool: initPool("localhost:6379", 16, 0, 300*time.Second),
		},
	}
	app := application{
		loop:    true,
		pool:    initPool("localhost:6379", 16, 0, 300*time.Second),
		service: service,
	}
	return app
}

func main() {

	app := initApp()
	app.showMenu()

	// app := application{
	// 	pool: initPool("localhost:6379", 16, 0, 300*time.Second),
	// }

	// // 测试添加一个 customer
	// customer := &model.Customer{
	// 	Id:     0,
	// 	Name:   "lightsaid",
	// 	Gender: "male",
	// 	Email:  "Ly@163.com",
	// }
	// dao := &dao.CustomerController{
	// 	Pool: app.pool,
	// }
	// err := dao.AddCustomer(customer)
	// if err != nil {
	// 	return
	// }
	// fmt.Println("添加成功")

	// // 取出一根连接池
	// conn := app.pool.Get()
	// // 延时关闭
	// defer conn.Close()

	// // 测试redis驱动是否成功
	// conn.Do("Set", "test", "测试")
	// res, err := redis.String(conn.Do("Get", "test"))
	// if err != nil {
	// 	fmt.Println("Get error:", err)
	// 	return
	// }
	// fmt.Println("res:", res)
}

package main

import (
	"fmt"
	"go-exercise/customer-manage/model"
	"go-exercise/customer-manage/service"

	"os"
)

type application struct {
	loop    bool
	key     int
	service *service.CustomerService
}

func showMenu(app *application) {
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
			app.add()
		case 2:
			fmt.Println("修改")
			app.update()
		case 3:
			app.delete()
		case 4:
			app.list()
		case 5:
			fmt.Println("退出")
			os.Exit(0)
		default:
			fmt.Print("输入有误，重新输入：")
		}
	}
}

func (app *application) add() {
	customer := &model.Customer{}
	fmt.Println("Name:")
	fmt.Scanln(&customer.Name)
	fmt.Println("Gender:")
	fmt.Scanln(&customer.Gender)
	fmt.Println("Email:")
	fmt.Scanln(&customer.Email)
	fmt.Println("customer:", customer)
	app.service.AddCustomer(customer)
	fmt.Println("Add Success")
}

func (app *application) list() {
	result := app.service.GetCustomer()
	fmt.Println(result)
}

func (app *application) update() {
	var id int
	name := ""
	gender := ""
	email := ""
	fmt.Println("Id:")
	fmt.Scanf("%d\n", &id)
	fmt.Println("Name:")
	fmt.Scanln(&name)
	fmt.Println("Gender:")
	fmt.Scanln(&gender)
	fmt.Println("Email:")
	fmt.Scanln(&email)
	b := app.service.Update(id, name, gender, email)
	if b {
		fmt.Println("Update Success")
	}
}

func (app *application) delete() {
	var id int
	fmt.Println("Id:")
	fmt.Scanf("%d\n", &id)
	b := app.service.Delete(id)
	if b {
		fmt.Println("Delete Success")
	}
}

func main() {
	server := &service.CustomerService{
		ID: 0,
	}
	app := &application{
		loop:    true,
		service: server,
	}
	showMenu(app)
}

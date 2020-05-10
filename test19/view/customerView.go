package main

import (
	"fmt"

	"github.com/unicornzyn/gotest19/model"

	"github.com/unicornzyn/gotest19/service"
)

//customerView view结构体
type customerView struct {
	key             string
	loop            bool
	customerService *service.CustomerService
}

//mainMenu 主菜单
func (cv *customerView) mainMenu() {
	for {
		fmt.Println("\n------------------客户信息管理软件------------------")
		fmt.Println("                    1 添 加 客 户")
		fmt.Println("                    2 修 改 客 户")
		fmt.Println("                    3 删 除 客 户")
		fmt.Println("                    4 客 户 列 表")
		fmt.Println("                    5 退 出 软 件")
		fmt.Print("请选择(1-5):")
		fmt.Scanln(&cv.key)
		switch cv.key {
		case "1":
			cv.add()
		case "2":
			cv.update()
		case "3":
			cv.delete()
		case "4":
			cv.list()
		case "5":
			cv.exit()
		default:
			fmt.Println("请输入正确的选项")
		}

		if !cv.loop {
			break
		}
	}
	fmt.Println("你退出了客户信息管理软件的使用")
}

//exit 退出软件
func (cv *customerView) exit() {
	fmt.Println("你确定要退出吗？ y/n")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Println("你的输入有误，请重新输入 y/n")
	}
	if choice == "y" {
		cv.loop = false
	}
}

//list 客户列表
func (cv *customerView) list() {
	customers := cv.customerService.List()
	fmt.Println("-----------------------------客户列表------------------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t\t邮箱")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("----------------------------客户列表完成---------------------------")
}

//add 添加客户
func (cv *customerView) add() {
	fmt.Println("-----------------------------添加客户------------------------------")
	fmt.Print("姓名:")
	name := ""
	fmt.Scanln(&name)
	fmt.Print("性别:")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Print("年龄:")
	age := 0
	fmt.Scanln(&age)
	fmt.Print("电话:")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Print("邮箱:")
	email := ""
	fmt.Scanln(&email)
	customer := model.NewCustomer2(name, gender, age, phone, email)
	cv.customerService.Add(customer)

	fmt.Println("----------------------------添加完成---------------------------")
}

func (cv *customerView) delete() {
	fmt.Println("-----------------------------删除客户------------------------------")
	fmt.Print("请选择删除客户编号(-1退出): ")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	fmt.Print("确认是否删除(y/n):")
	choice := ""
	fmt.Scanln(&choice)
	if choice == "Y" || choice == "y" {
		if cv.customerService.Delete(id) {
			fmt.Println("-----------------------------删除完成------------------------------")
		} else {
			fmt.Println("----------------------删除失败。输入的id不存在------------------------")
		}
	}

}

func (cv *customerView) update() {
	fmt.Println("-----------------------------修改客户------------------------------")
	fmt.Print("请选择修改客户编号(-1退出): ")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	index := cv.customerService.FindByID(id)
	if index == -1 {
		fmt.Println("--------------------------客户不存在----------------------------")
		return
	}
	c := cv.customerService.List()[index]
	fmt.Printf("姓名(%v):", c.Name)
	name := ""
	fmt.Scanln(&name)
	if name == "" {
		name = c.Name
	}
	fmt.Printf("性别(%v):", c.Gender)
	gender := ""
	fmt.Scanln(&gender)
	if gender == "" {
		gender = c.Gender
	}
	fmt.Printf("年龄(%v):", c.Age)
	age := 0
	fmt.Scanln(&age)
	if age == 0 {
		age = c.Age
	}
	fmt.Printf("电话(%v):", c.Phone)
	phone := ""
	fmt.Scanln(&phone)
	if phone == "" {
		phone = c.Phone
	}
	fmt.Printf("邮箱(%v):", c.Email)
	email := ""
	fmt.Scanln(&email)
	if email == "" {
		email = c.Email
	}
	customer := model.NewCustomer(c.ID, name, gender, age, phone, email)
	cv.customerService.Update(customer)

	fmt.Println("-----------------------------修改完成------------------------------")
}

func main() {
	cv := customerView{
		key:             "",
		loop:            true,
		customerService: service.NewCustomerService(),
	}
	cv.mainMenu()
}

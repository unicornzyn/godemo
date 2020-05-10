package utils

import "fmt"

// FamilyAccount 家庭收支软件结构体
type FamilyAccount struct {
	// 用户输入选项
	key string
	//循环标志
	loop bool
	//账户余额
	balance float64
	//每次收支的金额
	money float64
	//每次收支的说明
	note string
	//标示是否有收支的行为
	flag bool
	//每次收支的详情
	details string
}

// NewFamilyAccount 工厂方法生成FamilyAccount实例
func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		key:     "",
		loop:    true,
		balance: 10000.0,
		money:   0.0,
		note:    "",
		flag:    false,
		details: "收支\t账户金额\t收支金额\t说    明",
	}
}

// MainMenu 显示主菜单
func (acc *FamilyAccount) MainMenu() {
	for {
		fmt.Println("\n------------------家庭收支记账软件------------------")
		fmt.Println("                    1 收支明细")
		fmt.Println("                    2 登记收入")
		fmt.Println("                    3 登记支出")
		fmt.Println("                    4 退出软件")
		fmt.Print("请选择(1-4):")
		fmt.Scanln(&acc.key)
		switch acc.key {
		case "1":
			acc.showDetails()
		case "2":
			acc.income()
		case "3":
			acc.pay()
		case "4":
			acc.exit()
		default:
			fmt.Println("请输入正确的选项")
		}

		if !acc.loop {
			break
		}
	}
	fmt.Println("你退出了家庭记账软件的使用")
}

// ShowDetails 收支明细
func (acc *FamilyAccount) showDetails() {
	fmt.Println("------------------当前收支明细记录------------------")
	if acc.flag {
		fmt.Println(acc.details)
	} else {
		fmt.Println("当前没有收支...来一笔吧!")
	}
}

// income 登记收入
func (acc *FamilyAccount) income() {
	fmt.Println("本次收入金额:")
	fmt.Scanln(&acc.money)
	acc.balance += acc.money
	fmt.Println("本次收入说明:")
	fmt.Scanln(&acc.note)
	acc.details += fmt.Sprintf("\n收入\t%v\t%v\t%v", acc.balance, acc.money, acc.note)
	acc.flag = true
}

// pay 登记支出
func (acc *FamilyAccount) pay() {
	fmt.Println("本次支出金额:")
	fmt.Scanln(&acc.money)
	if acc.money > acc.balance {
		fmt.Println("支出的余额不足")
		return
	}
	acc.balance -= acc.money
	fmt.Println("本次支出说明:")
	fmt.Scanln(&acc.note)
	acc.details += fmt.Sprintf("\n支出\t%v\t%v\t%v", acc.balance, acc.money, acc.note)
	acc.flag = true
}

// exit 退出系统
func (acc *FamilyAccount) exit() {
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
		acc.loop = false
	}
}

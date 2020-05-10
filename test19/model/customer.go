package model

import (
	"fmt"
)

// Customer 客户信息结构体
type Customer struct {
	ID     int
	Name   string
	Gender string
	Age    int
	Phone  string
	Email  string
}

// NewCustomer 使用一个工厂模式 返回Cutomer的实例
func NewCustomer(id int, name string, gender string, age int, phone string, email string) Customer {
	return Customer{
		ID:     id,
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

// NewCustomer2 不带id的
func NewCustomer2(name string, gender string, age int, phone string, email string) Customer {
	return Customer{
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

//GetInfo 返回客户的信息
func (c Customer) GetInfo() string {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v", c.ID, c.Name, c.Gender,
		c.Age, c.Phone, c.Email)
	return info
}

package service

import (
	"github.com/unicornzyn/gotest19/model"
)

//CustomerService 完成对Customer的操作 包括增删改查
type CustomerService struct {
	customers []model.Customer
	//当前切片含有的客户数 该字段可以作为新客户的id+1
	customerNum int
}

//NewCustomerService 生成一个CustomerService实例
func NewCustomerService() *CustomerService {
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1, "张三", "男", 25, "133333333333", "zhangsan@sina.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}

//List 客户列表切片
func (cv *CustomerService) List() []model.Customer {
	return cv.customers
}

//Add 添加客户
func (cv *CustomerService) Add(c model.Customer) bool {
	cv.customerNum++
	c.ID = cv.customerNum
	cv.customers = append(cv.customers, c)
	return true
}

//FindByID 根据id查找客户下标 如果不存在返回-1
func (cv *CustomerService) FindByID(id int) int {
	index := -1
	for i := 0; i < len(cv.customers); i++ {
		if cv.customers[i].ID == id {
			index = i
			break
		}
	}
	return index
}

//Delete 根据id删除切片中的客户
func (cv *CustomerService) Delete(id int) bool {
	index := cv.FindByID(id)
	if index == -1 {
		return false
	}
	//删除切片中的元素
	cv.customers = append(cv.customers[:index], cv.customers[index+1:]...)
	return true
}

//Update 更新客户信息
func (cv *CustomerService) Update(c model.Customer) bool {
	index := cv.FindByID(c.ID)
	if index == -1 {
		return false
	}
	cv.customers[index].Name = c.Name
	cv.customers[index].Gender = c.Gender
	cv.customers[index].Age = c.Age
	cv.customers[index].Phone = c.Phone
	cv.customers[index].Email = c.Email
	return true
}

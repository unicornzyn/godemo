package main

import (
	"fmt"

	"github.com/unicornzyn/gotest15/model"
)

func main() {
	stu := model.NewStudent("jack", 100.0)
	fmt.Println(*stu)
}

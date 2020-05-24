package model

import (
	"errors"
)

//根据业务逻辑需要 自定义错误

var (
	//ErrorUserNotExists err
	ErrorUserNotExists = errors.New("用户不存在")
	//ErrorUserExists err
	ErrorUserExists = errors.New("用户已存在")
	//ErrorUserPwd err
	ErrorUserPwd = errors.New("密码错误")
)

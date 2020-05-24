package model

import (
	"encoding/json"
	"fmt"

	"github.com/garyburd/redigo/redis"
)

// 定义全局变量 在服务器启动时初始化
var (
	MyUserDao *UserDao
)

// UserDao 对User的操作
type UserDao struct {
	pool *redis.Pool
}

// UserDaoNew 使用工厂模式创建UserDao的实例
func UserDaoNew(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return
}

// getUserByID 根据id查询user
func (dao *UserDao) getUserByID(conn redis.Conn, id int) (user *User, err error) {
	res, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		if err == redis.ErrNil {
			err = ErrorUserNotExists
		}
		return
	}
	user = &User{}
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(res), user) err=", err)
	}
	return
}

// Login 登录逻辑
func (dao *UserDao) Login(userID int, userPwd string) (user *User, err error) {
	conn := dao.pool.Get()
	defer conn.Close()
	user, err = dao.getUserByID(conn, userID)
	if err != nil {
		return
	}
	if user.UserPwd != userPwd {
		err = ErrorUserPwd
		return
	}
	return
}

// Register 注册
func (dao *UserDao) Register(user *User) (err error) {
	conn := dao.pool.Get()
	defer conn.Close()
	_, err = dao.getUserByID(conn, user.UserID)
	if err == nil {
		err = ErrorUserExists
		return
	}

	data, err := json.Marshal(user)
	if err != nil {
		return
	}
	fmt.Println(user)
	_, err = conn.Do("HSet", "users", user.UserID, string(data))
	return
}

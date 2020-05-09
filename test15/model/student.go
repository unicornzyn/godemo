package model

// student 定义一个model
type student struct {
	Name  string
	Score float64
}

// NewStudent student第一个字母s小写 其他包无法访问 可以使用工厂模式生成student
func NewStudent(n string, s float64) *student {
	return &student{
		Name:  n,
		Score: s,
	}
}

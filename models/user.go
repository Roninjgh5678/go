package models

//只能大写被外部引用
type User struct {
	Id       int
	Username string
	Age      int
	Email    string
	AddTime  int
}

//表示配置操作数据库的表名称
func (User) TableName() string {

	return "users"
}

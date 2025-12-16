package models

//只能大写被外部引用
type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	AddTime  int    `json:"addTime"`
}

//表示配置操作数据库的表名称
func (User) TableName() string {

	return "users"
}

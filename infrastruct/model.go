// 基础设施层，相当于数据持久化层
package infrastruct

// User 用户model
type User struct {
	UserId      int `json:userId,omitempty`
	UserName    string
	BirthOfDate string
	Address     string
	Description string
	CreateAt    string
}

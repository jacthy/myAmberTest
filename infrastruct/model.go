// 基础设施层，相当于数据持久化层
package infrastruct

import "time"

// User 用户model
type User struct {
	UserId      string `json:userId,omitempty`
	UserName    string
	BirthOfDate *time.Time
	Address     string
	Description string
	CreateAt    *time.Time
	//Origin *User
}

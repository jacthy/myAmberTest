// 数据持久化层
package infrastruct

import "time"

// User 用户model
type User struct {
	userId      string
	userName    string
	birthOfDate *time.Time
	address     string
	description string
	createAt    *time.Time
	//Origin *User
}

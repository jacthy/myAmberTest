package repo

import (
	"github.com/liaojuntao/infrastruct"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

type userModel struct {
	userId string
	userName string
	birthOfDate *time.Time
	address string
	description string
	createAt *time.Time
}

// DefaultUserRepo 默认用户实体的存储服务实例
var DefaultUserRepo *UserRepo

func init() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// 这里的数据库初始化应该独立DB操作，这里为了简化demo所以耦合在这里
	db.AutoMigrate(&userModel{})
	DefaultUserRepo = &UserRepo{
		DB: db,
	}
}

type UserRepo struct {
	DB *gorm.DB
}

// NewUserRepo 返回sqlite的仓储存储服务实例（单例饿汉模式）
func NewUserRepo() infrastruct.UserRepo {
	return DefaultUserRepo
}

func (u UserRepo) Create(user *infrastruct.User) error {
	panic("implement me")
}

func (u UserRepo) Update(user *infrastruct.User) error {
	panic("implement me")
}

func (u UserRepo) Delete(user *infrastruct.User) error {
	panic("implement me")
}

func (u UserRepo) GetById(id int) *infrastruct.User {
	panic("implement me")
}

package repo

import (
	"github.com/liaojuntao/infrastruct"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

type user struct {
	UserId      string `gorm:"primary_key"`
	UserName    string
	BirthOfDate *time.Time
	Address     string
	Description string
	CreateAt    *time.Time
}

// defaultUserRepo 默认用户实体的存储服务实例
var defaultUserRepo *UserRepo

func init() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 这里的数据库初始化应该独立DB操作，这里为了简化demo所以耦合在这里
	if !db.Migrator().HasTable(&user{}) {
		db.AutoMigrate(&user{})
	}
	defaultUserRepo = &UserRepo{
		DB: db,
	}
}

type UserRepo struct {
	DB *gorm.DB
}

// GetUserRepo 返回sqlite的仓储存储服务实例（单例饿汉模式）
func GetUserRepo() infrastruct.UserRepo {
	return defaultUserRepo
}

// Create 将新建对象进行持久化
func (u *UserRepo) Create(user *infrastruct.User) error {
	return u.DB.Create(user).Error
}

func (u *UserRepo) DeleteById(id int) error {
	return u.DB.Delete(&user{}, "user_id", id).Error
}

func (u *UserRepo) GetByUserName(userName string) (*infrastruct.User, error) {
	user, err := u.getByUserName(userName)
	if err != gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepo) getByUserName(userName string) (*infrastruct.User, error) {
	var user infrastruct.User
	err := u.DB.First(&user, "user_name", userName).Error
	return &user, err
}

func (u *UserRepo) NotExistByName(userName string) (bool, error) {
	_, err := u.getByUserName(userName)
	if err == gorm.ErrRecordNotFound {
		return true, nil
	}
	return false, err
}

func (u *UserRepo) Update(user *infrastruct.User) error {
	panic("implement me")
}

package repo

import (
	"github.com/liaojuntao/infrastruct"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

// defaultUserRepo 默认用户实体的存储服务实例
var defaultUserRepo *UserRepo

func init() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// 为便于检查，现设置debug 模式
	db.Debug()
	// 这里的数据库初始化应该独立DB操作，这里为了简化demo所以耦合在这里
	if !db.Migrator().HasTable(&user{}) {
		db.AutoMigrate(&user{})
	}
	defaultUserRepo = &UserRepo{
		db: db,
	}
}

type UserRepo struct {
	db *gorm.DB
}

// GetUserRepo 返回sqlite的仓储存储服务实例（单例饿汉模式）
func GetUserRepo() infrastruct.UserRepo {
	return defaultUserRepo
}

// Create 将新建对象进行持久化
func (u *UserRepo) Create(user *infrastruct.User) error {
	model := toSqliteModel(user)
	currentTime := time.Now()
	model.CreateAt = &currentTime
	return u.db.Create(model).Error
}

func (u *UserRepo) DeleteById(id int) error {
	return u.db.Delete(&user{}, "user_id", id).Error
}

func (u *UserRepo) GetByUserName(userName string) (*infrastruct.User, error) {
	user, err := u.getByUserName(userName)
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepo) getByUserName(userName string) (*infrastruct.User, error) {
	var user user
	err := u.db.First(&user, "user_name", userName).Error
	return toUser(&user), err
}

func (u *UserRepo) NotExistByName(userName string) (bool, error) {
	_, err := u.getByUserName(userName)
	if err == gorm.ErrRecordNotFound {
		return true, nil
	}
	return false, err
}

func (u *UserRepo) Update(user *infrastruct.User) error {
	modelUser := toSqliteModel(user)
	return u.db.Model(&modelUser).Updates(modelUser).Error
}

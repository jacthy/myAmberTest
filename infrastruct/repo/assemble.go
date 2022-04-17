package repo

import (
	"github.com/liaojuntao/infrastruct"
	"time"
)

const defaultTimeFormat = "2006-01-02 15:04:05"

func toSqliteModel(userDto *infrastruct.User) *user {
	if userDto == nil {
		return nil
	}
	return &user{
		UserId:      userDto.UserId,
		UserName:    userDto.UserName,
		BirthOfDate: userDto.BirthOfDate,
		Address:     userDto.Address,
		Description: userDto.Description,
	}
}

func toUser(model *user) *infrastruct.User {
	if model == nil {
		return nil
	}
	return &infrastruct.User{
		UserId:      model.UserId,
		UserName:    model.UserName,
		BirthOfDate: model.BirthOfDate,
		Address:     model.Address,
		Description: model.Description,
		CreateAt:    model.CreateAt.Format(defaultTimeFormat),
	}
}

type user struct {
	UserId      int `gorm:"primary_key"`
	UserName    string
	BirthOfDate string
	Address     string
	Description string
	CreateAt    *time.Time
}

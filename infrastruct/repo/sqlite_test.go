package repo

import (
	"github.com/liaojuntao/infrastruct"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

// 同时测试 Create,GetByUserName,DeleteById，Update 因为刚好构成setup->exec->clear的关系
func TestCreate_GetByUserName_Update_DeleteById(t *testing.T) {
	defer func() {
		// 防止产生脏数据影响其他测试用例
		defaultUserRepo.db.Exec("DELETE FROM  users")
	}()
	Convey("testing Create,GetByUserName,DeleteById", t, func() {
		// clear dirty data 防止脏数据污染
		err := defaultUserRepo.db.Exec("DELETE FROM  users").Error
		if err != nil {
			t.Fatal("repo run error: ", err.Error())
			return
		}
		testingUser := infrastruct.User{
			UserName:    "user1",
			BirthOfDate: "2013.09.09",
			Address:     "Address",
			Description: "Description",
		}
		// 新建用户实体
		err = defaultUserRepo.Create(&testingUser)
		So(err, ShouldBeNil)
		// 获取用户实体
		data, err2 := defaultUserRepo.GetByUserName(testingUser.UserName)
		So(err2, ShouldBeNil)
		So(data.UserName, ShouldEqual, testingUser.UserName)
		So(data.BirthOfDate, ShouldEqual, testingUser.BirthOfDate)
		So(data.Address, ShouldEqual, testingUser.Address)
		So(data.Description, ShouldEqual, testingUser.Description)
		So(data.CreateAt, ShouldEqual, time.Now().Format(defaultTimeFormat))

		// 更新用户实体
		testingUser.UserId = data.UserId
		testingUser.UserName = "user2"
		testingUser.BirthOfDate = "2013.09.02"
		testingUser.Address = "Address2"
		testingUser.Description = "Description2"
		testingUser.CreateAt = defaultTimeFormat
		defaultUserRepo.Update(&testingUser)
		newData, err := defaultUserRepo.GetByUserName(testingUser.UserName)
		So(err, ShouldBeNil)
		So(newData.UserName, ShouldEqual, "user2")
		So(newData.BirthOfDate, ShouldEqual, "2013.09.02")
		So(newData.Address, ShouldEqual, "Address2")
		So(newData.Description, ShouldEqual, "Description2")
		So(newData.CreateAt, ShouldEqual, time.Now().Format(defaultTimeFormat))

		// 删除客户实体
		err = defaultUserRepo.DeleteById(data.UserId)
		So(err, ShouldBeNil)
	})
}

func TestGetByUserId(t *testing.T) {
	defer func() {
		// 防止产生脏数据影响其他测试用例
		defaultUserRepo.db.Exec("DELETE FROM  users")
	}()
	// setup
	err := defaultUserRepo.db.Exec("DELETE FROM  users;INSERT INTO \"users\" VALUES (111, 'user111', '2013.09.02', 'Address2', 'Description2', '2022-04-17 23:21:13.27929+08:00');").Error
	if err != nil {
		t.Fatal("repo run error: ", err.Error())
		return
	}
	Convey("testing GetByUserById", t, func() {
		Convey("testing GetByUserById when succeed", func() {
			user_, err := defaultUserRepo.GetByUserId(111)
			So(err, ShouldBeNil)
			So(user_.UserId, ShouldEqual, 111)
			So(user_.UserName, ShouldEqual, "user111")
		})

		Convey("testing GetByUserById when not exit", func() {
			user_, err := defaultUserRepo.GetByUserId(333)
			So(err, ShouldBeNil)
			So(user_, ShouldBeNil)
		})
	})



}

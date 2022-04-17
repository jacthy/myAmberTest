package repo

import (
	"github.com/liaojuntao/infrastruct"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

// 同时测试 Create,GetByUserName,DeleteById，因为刚好构成setup->exec->clear的关系
func TestCreate_GetByUserName_DeleteById(t *testing.T) {
	defer func() {
		// 退出测试后删除脏数据
		defaultUserRepo.db.Exec("DELETE FROM  users")
	}()
	Convey("testing Create,GetByUserName,DeleteById", t, func() {
		// clear dirty data
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


package router

import (
	"errors"
	. "github.com/agiledragon/gomonkey"
	"github.com/liaojuntao/controller"
	"github.com/liaojuntao/infrastruct"
	. "github.com/smartystreets/goconvey/convey"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

// 测试创建用户
func Test_CreateUserRouter_Handler(t *testing.T) {
	patch := NewPatches()
	defer patch.Reset()

	Convey("testing CreateUserRouter handler", t, func() {

		Convey("testing CreateUserRouter when success", func() {
			// 打桩controller层，router.handler的业务应与controller层解耦，所以测试用例也应该解藕
			patch.ApplyMethod(reflect.TypeOf(&controller.UserCtl{}), "CreateUser",
				func(_ *controller.UserCtl, _ *infrastruct.User) error {
					return nil
				})

			userRou := CreateUserRouter()
			postData := "{\"userName\":\"用户1\",\"birthOfDate\":\"2021-02-09\",\"address\":\"广州\",\"description\":\"描述1\"}"
			respMock := httptest.NewRecorder()
			reqMock := new(http.Request)
			reqMock.Body = ioutil.NopCloser(strings.NewReader(postData))
			userRou.GetHandler()(respMock, reqMock)
			So(respMock.Body.String(), ShouldEqual, "{\"status\":2000,\"data\":\"succeed\"}")
			So(respMock.Code, ShouldEqual, 200)
		})

		Convey("testing CreateUserRouter handler when fail with wrong param", func() {
			userRou := CreateUserRouter()
			postData := "{\"userName\":1000,\"birthOfDate\":\"2021-02-09\",\"address\":\"广州\",\"description\":\"描述1\"}"
			respMock := httptest.NewRecorder()
			reqMock := new(http.Request)
			reqMock.Body = ioutil.NopCloser(strings.NewReader(postData))
			userRou.GetHandler()(respMock, reqMock)
			So(respMock.Body.String(), ShouldEqual, "{\"errCode\":4001,\"errMessage\":\"参数校验错误\"}")
			So(respMock.Code, ShouldEqual, 200)
		})

		Convey("testing CreateUserRouter handler when fail with repeat user name", func() {
			patch.Reset()
			// 打桩controller层，router.handler的业务应与controller层解耦，所以测试用例也应该解藕
			patch.ApplyMethod(reflect.TypeOf(&controller.UserCtl{}), "CreateUser",
				func(_ *controller.UserCtl, _ *infrastruct.User) error {
					return errors.New("该用户已存在")
				})
			userRou := CreateUserRouter()
			postData := "{\"userName\":\"用户1\",\"birthOfDate\":\"2021-02-09\",\"address\":\"广州\",\"description\":\"描述1\"}"
			respMock := httptest.NewRecorder()
			reqMock := new(http.Request)
			reqMock.Body = ioutil.NopCloser(strings.NewReader(postData))
			userRou.GetHandler()(respMock, reqMock)
			So(respMock.Body.String(), ShouldEqual, "{\"errCode\":4002,\"errMessage\":\"该用户已存在\"}")
			So(respMock.Code, ShouldEqual, 200)
		})
	})
}

// 测试更新用户
func Test_UpdateUserRouter_Handler(t *testing.T) {
	patch := NewPatches()
	defer patch.Reset()

	Convey("testing UpdateUserRouter handler", t, func() {

		Convey("testing UpdateUserRouter when success", func() {
			// 打桩controller层，router.handler的业务应与controller层解耦，所以测试用例也应该解藕
			patch.ApplyMethod(reflect.TypeOf(&controller.UserCtl{}), "UpdateUser",
				func(_ *controller.UserCtl, _ *infrastruct.User) error {
					return nil
				})

			userUpdateRou := UpdateUserRouter()
			postData := "{\"userId\":1,\"userName\":\"用户1\",\"birthOfDate\":\"2021-02-09\",\"address\":\"广州\",\"description\":\"描述1\"}"
			respMock := httptest.NewRecorder()
			reqMock := new(http.Request)
			reqMock.Body = ioutil.NopCloser(strings.NewReader(postData))
			userUpdateRou.GetHandler()(respMock, reqMock)
			So(respMock.Body.String(), ShouldEqual, "{\"status\":2000,\"data\":\"succeed\"}")
			So(respMock.Code, ShouldEqual, 200)
		})

		Convey("testing UpdateUserRouter handler when fail with wrong param", func() {
			userUpdateRou := UpdateUserRouter()
			postData := "{\"userName\":1000,\"birthOfDate\":\"2021-02-09\",\"address\":\"广州\",\"description\":\"描述1\"}"
			respMock := httptest.NewRecorder()
			reqMock := new(http.Request)
			reqMock.Body = ioutil.NopCloser(strings.NewReader(postData))
			userUpdateRou.GetHandler()(respMock, reqMock)
			So(respMock.Body.String(), ShouldEqual, "{\"errCode\":4001,\"errMessage\":\"参数校验错误\"}")
		})

		Convey("testing CreateUserRouter handler when fail with repeat user name", func() {
			patch.Reset()
			// 打桩controller层，router.handler的业务应与controller层解耦，所以测试用例也应该解藕
			patch.ApplyMethod(reflect.TypeOf(&controller.UserCtl{}), "UpdateUser",
				func(_ *controller.UserCtl, _ *infrastruct.User) error {
					return errors.New("该用户已存在")
				})
			userUpdateRou := UpdateUserRouter()
			postData := "{\"userName\":\"用户1\",\"birthOfDate\":\"2021-02-09\",\"address\":\"广州\",\"description\":\"描述1\"}"
			respMock := httptest.NewRecorder()
			reqMock := new(http.Request)
			reqMock.Body = ioutil.NopCloser(strings.NewReader(postData))
			userUpdateRou.GetHandler()(respMock, reqMock)
			So(respMock.Body.String(), ShouldEqual, "{\"errCode\":4002,\"errMessage\":\"该用户已存在\"}")
		})
	})
}

// 测试根据id获取用户
func Test_GetUserByIdRouter_Handler(t *testing.T) {
	patch := NewPatches()
	defer patch.Reset()
	userData := "{\"userId\":1,\"userName\":\"用户1\",\"birthOfDate\":\"2021-02-09\",\"address\":\"广州\",\"description\":\"描述1\"}"
	Convey("testing getById handler", t, func() {

		Convey("testing getById handler when success", func() {
			// 打桩controller层，router.handler的业务应与controller层解耦，所以测试用例也应该解藕
			patch.ApplyMethod(reflect.TypeOf(&controller.UserCtl{}), "GetUserById",
				func(_ *controller.UserCtl, _ int) (string, error) {
					return userData, nil
				})

			userGetByIdRou := GetByIdRouter()
			respMock := httptest.NewRecorder()
			reqMock := httptest.NewRequest("GET", "http://www.baidu.com/user/getById?userId=1",
				new(strings.Reader))
			userGetByIdRou.GetHandler()(respMock, reqMock)
			So(respMock.Body.String(), ShouldStartWith, "{\"status\":2000,\"data\":")
			So(respMock.Code, ShouldEqual, 200)
		})

		Convey("testing getById handler when fail with wrong param", func() {
			userGetByIdRou := GetByIdRouter()
			respMock := httptest.NewRecorder()
			reqMock := httptest.NewRequest("GET", "http://www.baidu.com/user/getById?userId=ksdjhf",
				new(strings.Reader))
			userGetByIdRou.GetHandler()(respMock, reqMock)
			So(respMock.Body.String(), ShouldEqual, "{\"errCode\":4001,\"errMessage\":\"参数校验错误\"}")
		})

		Convey("testing getById handler when empty with not exist userId", func() {
			patch.Reset()
			// 打桩controller层，router.handler的业务应与controller层解耦，所以测试用例也应该解藕
			patch.ApplyMethod(reflect.TypeOf(&controller.UserCtl{}), "GetUserById",
				func(_ *controller.UserCtl, _ int) (string, error) {
					return "", nil
				})
			userGetByIdRou := GetByIdRouter()
			respMock := httptest.NewRecorder()
			reqMock := httptest.NewRequest("GET", "http://www.baidu.com/user/getById?userId=222",
				new(strings.Reader))
			userGetByIdRou.GetHandler()(respMock, reqMock)
			So(respMock.Body.String(), ShouldEqual, "{\"status\":2000,\"data\":\"\"}")
		})
	})
}

// 测试根据id删除用户， UserController.DeleteUserById的方法被编译器内联了，所以这个测试用例需要禁用内联优化才可以成功
func Test_DeleteUserByIdRouter_Handler(t *testing.T) {
	patch := NewPatches()
	defer patch.Reset()

	// 这里的mock需要禁用内联优化才可以成功
	// 打桩controller层，router.handler的业务应与controller层解耦，所以测试用例也应该解藕
	patch.ApplyMethod(reflect.TypeOf(controller.NewUserController(nil)), "DeleteUserById",
		func(_ *controller.UserCtl, mockId int) error {
			println("heated it")
			if mockId == 222 {
				return errors.New("controller err")
			}
			return nil
		})

	Convey("testing DeleteUserById handler", t, func() {

		Convey("testing DeleteUserById handler when success", func() {

			userDeleteByIdRou := DeleteByIdRouter()
			respMock := httptest.NewRecorder()
			reqMock := httptest.NewRequest("DELETE", "http://www.baidu.com/user/getById?userId=1",
				new(strings.Reader))
			userDeleteByIdRou.GetHandler()(respMock, reqMock)
			So(respMock.Body.String(), ShouldStartWith, "{\"status\":2000,\"data\":")
			So(respMock.Code, ShouldEqual, 200)
		})

		Convey("testing DeleteUserById handler when fail with wrong param", func() {
			userDeleteByIdRou := DeleteByIdRouter()
			respMock := httptest.NewRecorder()
			reqMock := httptest.NewRequest("GET", "http://www.baidu.com/user/getById?userId=ksdjhf",
				new(strings.Reader))
			userDeleteByIdRou.GetHandler()(respMock, reqMock)
			So(respMock.Body.String(), ShouldEqual, "{\"errCode\":4001,\"errMessage\":\"参数校验错误\"}")
		})

		Convey("testing DeleteUserById handler when controller err", func() {

			userDeleteByIdRou := DeleteByIdRouter()
			respMock := httptest.NewRecorder()
			reqMock := httptest.NewRequest("GET", "http://www.baidu.com/user/getById?userId=222",
				new(strings.Reader))
			userDeleteByIdRou.GetHandler()(respMock, reqMock)
			So(respMock.Body.String(), ShouldEqual, "{\"errCode\":4002,\"errMessage\":\"controller err\"}")
		})
	})
}

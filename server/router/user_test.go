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

			userRou := GetUserRouter().CreateUserRouter()
			postData := "{\"userName\":\"用户1\",\"birthOfDate\":\"2021-02-09\",\"address\":\"广州\",\"description\":\"描述1\"}"
			respMock := httptest.NewRecorder()
			reqMock := new(http.Request)
			reqMock.Body = ioutil.NopCloser(strings.NewReader(postData))
			userRou.GetHandler()(respMock, reqMock)
			So(respMock.Body.String(), ShouldEqual, "{\"status\":2000,\"data\":\"succeed\"}")
			So(respMock.Code, ShouldEqual, 200)
		})

		Convey("testing CreateUserRouter handler when fail with wrong param", func() {
			userRou := GetUserRouter().CreateUserRouter()
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
			userRou := GetUserRouter().CreateUserRouter()
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

			userUpdateRou := GetUserRouter().UpdateUserRouter()
			postData := "{\"userId\":1,\"userName\":\"用户1\",\"birthOfDate\":\"2021-02-09\",\"address\":\"广州\",\"description\":\"描述1\"}"
			respMock := httptest.NewRecorder()
			reqMock := new(http.Request)
			reqMock.Body = ioutil.NopCloser(strings.NewReader(postData))
			userUpdateRou.GetHandler()(respMock, reqMock)
			So(respMock.Body.String(), ShouldEqual, "{\"status\":2000,\"data\":\"succeed\"}")
			So(respMock.Code, ShouldEqual, 200)
		})

		Convey("testing UpdateUserRouter handler when fail with wrong param", func() {
			userUpdateRou := GetUserRouter().UpdateUserRouter()
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
			userUpdateRou := GetUserRouter().UpdateUserRouter()
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

			userGetByIdRou := GetUserRouter().GetByIdRouter()
			respMock := httptest.NewRecorder()
			reqMock := httptest.NewRequest("GET", "http://www.baidu.com/user/getById?userId=1",
				new(strings.Reader))
			userGetByIdRou.GetHandler()(respMock, reqMock)
			So(respMock.Body.String(), ShouldStartWith, "{\"status\":2000,\"data\":")
			So(respMock.Code, ShouldEqual, 200)
		})

		Convey("testing getById handler when fail with wrong param", func() {
			userGetByIdRou := GetUserRouter().GetByIdRouter()
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
			userGetByIdRou := GetUserRouter().GetByIdRouter()
			respMock := httptest.NewRecorder()
			reqMock := httptest.NewRequest("GET", "http://www.baidu.com/user/getById?userId=222",
				new(strings.Reader))
			userGetByIdRou.GetHandler()(respMock, reqMock)
			So(respMock.Body.String(), ShouldEqual, "{\"status\":2000,\"data\":\"\"}")
		})
	})
}

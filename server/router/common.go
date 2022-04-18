package router

import (
	"encoding/json"
	"github.com/liaojuntao/infrastruct"
	"io/ioutil"
	"net/http"
	"strconv"
)

const (
	statusOK = 2000 // 成功响应
	paramErr = 4001 // 参数校验错误
	optErr   = 4002 // 业务操作错误
)

// ErrResp 错误时的响应
type ErrResp struct {
	ErrCode    int    `json:"errCode"`    // 错误响应码
	ErrMessage string `json:"errMessage"` // 错误信息
}

// SuccessResp 成功时的响应
type SuccessResp struct {
	Status int    `json:"status"` // 响应码
	Data   string `json:"data"`   // json数据/string
}

// setParamErr 设置参数错误
func setParamErr(resp http.ResponseWriter) {
	err := ErrResp{
		ErrCode:    paramErr,
		ErrMessage: "参数校验错误",
	}
	message, e := json.Marshal(err)
	if e != nil {
		println(e.Error()) // 应通过日志记录，这里简化
		// 由于此处marshal错误，也需要writeHeader,所以不直接return
	}
	resp.Write(message)
}

// setErrResp 设置业务执行错误
func setErrResp(resp http.ResponseWriter, errorMessage string) {
	err := ErrResp{
		ErrCode:    optErr,
		ErrMessage: errorMessage,
	}
	message, e := json.Marshal(err)
	if e != nil {
		println(e.Error()) // 同setParamErr中marshal的错误处理
	}
	resp.Write(message)
	resp.WriteHeader(http.StatusNotAcceptable)
}

// setSuccessResp 成功响应信息
// message 参数可以是字符串，可以是数据结构
func setSuccessResp(resp http.ResponseWriter, message string) {
	reap := SuccessResp{
		Status: statusOK,
		Data:   message,
	}
	msg, e := json.Marshal(reap)
	if e != nil {
		println(e.Error()) // 同setParamErr中marshal的错误处理
	}
	resp.Write(msg)
	resp.WriteHeader(http.StatusOK)
}

// getUserModelFromReqBody 从http body取参数
func getUserModelFromReqBody(req *http.Request) (*infrastruct.User, error) {
	userModel := new(infrastruct.User)
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(result, userModel); err != nil {
		return nil, err
	}
	return userModel, nil
}

func getIdFromReqHeader(req *http.Request) (int, error) {
	return strconv.Atoi(req.URL.Query().Get("userId"))
}

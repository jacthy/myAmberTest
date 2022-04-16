package router

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ErrResp 错误时的响应
type ErrResp struct {
	errCode    int
	errMessage string
}

// setParamErr 设置参数错误
func setParamErr(resp http.ResponseWriter) {
	err := ErrResp{
		errCode:    paramErr,
		errMessage: "参数校验错误",
	}
	message, e := json.Marshal(err)
	if e != nil {
		println(e.Error()) // 应通过日志记录，这里简化
		// 由于此处marshal错误，也需要writeHeader,所以不直接return
	}
	fmt.Printf("%v\n", string(message))
	resp.Write(message)
	resp.WriteHeader(http.StatusNotAcceptable)
}

// setErrResp 设置业务执行错误
func setErrResp(resp http.ResponseWriter, errorMessage string) {
	err := ErrResp{
		errCode:    optErr,
		errMessage: errorMessage,
	}
	message, e := json.Marshal(err)
	if e != nil {
		println(e.Error()) // 同setParamErr中marshal的错误处理
	}
	resp.Write(message)
	resp.WriteHeader(http.StatusNotAcceptable)
}

// setSuccessResp 成功响应信息
func setSuccessResp(resp http.ResponseWriter, message interface{}) {
	if message != nil {
		msg, e := json.Marshal(message)
		if e != nil {
			println(e.Error()) // 同setParamErr中marshal的错误处理
		}
		resp.Write(msg)
	}
	resp.WriteHeader(http.StatusOK)
}

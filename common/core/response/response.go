package response

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"t3-zero/common/core/errorx"
)

type Body struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type Empty struct {
}

const (
	ERROR   = 1
	SUCCESS = 0
)

/**
 * @description: 通用请求返回
 * @param {http.ResponseWriter} w
 * @param {interface{}} resp 写入json消息
 * @param {error} err  错误信息
 * @return {*}
 */
func Response(w http.ResponseWriter, resp interface{}, err error) {
	var body Body

	switch e := err.(type) {
	case nil:
		body.Msg = "OK"
		// 如果返回值为nil,返回空类型对象
		// 注意如果返回值类型并非显式interface{} 此时会判断为非空
		if resp == nil {
			body.Data = Empty{}
		} else {
			body.Data = resp
		}
	case *errorx.CodeError:
		body.Code = e.Data().Code
		body.Msg = e.Data().Msg
		body.Data = Empty{}
	default:
		body.Code = ERROR
		body.Msg = err.Error()
		body.Data = Empty{}
	}
	// 返回错误信息打印链路日志
	if err != nil {
		logx.ErrorStack(err)
	}
	httpx.OkJson(w, body)
}

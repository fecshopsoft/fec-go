package util

import (
	"errors"
)

const (
	ResultStatusNeedPermission = 40000
	ResultStatusNeedLogin      = 30000
	ResultStatusSuccess        = 20000
	ResultStatusFail           = 10000
)

var ErrNeedMiLogin = errors.New("Please login to your account first")
var ErrNeedMiPermission = errors.New("You do not have permission to operate this resource")

type ResultVO struct {
	Code int         `form:"code" json:"code"`
	Msg    string      `form:"msg" json:"msg"`
	Data interface{} `form:"data" json:"data"`
}

type PageVO struct {
	TargetPage int         `form:"targetPage" json:"targetPage"`
	PageSize   int         `form:"pageSize" json:"pageSize"`
	Total      int         `form:"total" json:"total"`
	TotalPage  int         `form:"totalPage" json:"totalPage"`
	Datas      interface{} `form:"datas" json:"datas"`
}

func BuildSuccessResult(data interface{}) *ResultVO {
	result := &ResultVO{ResultStatusSuccess, "", data}
	return result
}

func BuildFailResult(msg string) *ResultVO {
	result := &ResultVO{ResultStatusFail, msg, nil}
	return result
}

func BuildNeedLoginResult() *ResultVO {
	result := &ResultVO{ResultStatusNeedLogin, ErrNeedMiLogin.Error(), nil}
	return result
}

func BuildNeedPermissionResult() *ResultVO {
	result := &ResultVO{ResultStatusNeedPermission, ErrNeedMiPermission.Error(), nil}
	return result
}



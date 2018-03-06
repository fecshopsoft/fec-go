package util

import (
	"errors"
)

const (
	ResultStatusNeedPermission = 3
	ResultStatusNeedLogin      = 2
	ResultStatusSuccess        = 1
	ResultStatusFail           = 0
)

var ErrNeedMiLogin = errors.New("请先登录再进行下一步操作")
var ErrNeedMiPermission = errors.New("没有足够权限进行下一步操作")

type ResultVO struct {
	Status int         `form:"status" json:"status"`
	Msg    string      `form:"msg" json:"msg"`
	Result interface{} `form:"result" json:"result"`
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



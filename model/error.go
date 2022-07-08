package model

import (
	"github.com/FengZhg/go_tools/errs"
)

const (
	errorHTTPResponseStatus = 259000 + iota
	errorScrapeTimeout
	errorConfigNotFound
	errorResponse
	errorUid
)

var (
	HTTPResponseStatusError = errs.NewError(errorHTTPResponseStatus, "HTTP请求相应状态错误")
	ScrapeTimeoutError      = errs.NewError(errorScrapeTimeout, "处理调度超时")
	ConfigNotFoundError     = errs.NewError(errorConfigNotFound, "配置文件缺失")
	ResponseError           = errs.NewError(errorResponse, "响应错误")
	UidError                = errs.NewError(errorUid, "用户id缺失")
)

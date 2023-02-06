package api

import "github.com/XCPCBoard/common/errors"

var (
	UserExistError = errors.NewError(212, "用户名或邮件已存在")
)

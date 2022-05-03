package errtype

import (
	"fmt"

	"github.com/fangzhixi/go-common/error/errcode"
)

//错误基类
type ErrorBase struct {
	LogId       string
	Code        int32
	OriginError error
}

func NewError(logId string, code int32, err error) ErrorBase {
	return ErrorBase{
		LogId:       logId,
		Code:        code,
		OriginError: err,
	}
}

func (e ErrorBase) Error() string {
	return fmt.Sprintf("%s - %s", e.LogId, errcode.ErrorMsg[e.Code])
}

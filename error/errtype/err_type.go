package errtype

import (
	"fmt"

	"github.com/fangzhixi/go-common/error/errcode"
)

//异常基类
type BaseError struct {
	LogId       string
	Code        int32
	OriginError error
}

func NewBaseError(logId string, code int32, err error) *BaseError {
	return &BaseError{
		LogId:       logId,
		Code:        code,
		OriginError: err,
	}
}

func (e *BaseError) Error() string {
	if e == nil {
		return ""
	}
	return fmt.Sprintf("%d:%s - %s", e.Code, errcode.GetErrMsg(e.Code), e.LogId)
}

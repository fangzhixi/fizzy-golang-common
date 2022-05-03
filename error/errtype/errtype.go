package errtype

import (
	"fmt"

	"github.com/fangzhixi/go-common/error/errcode"
)

//错误基类
type BusinessError struct {
	LogId       string
	Code        int32
	OriginError error
}

func NewBusError(logId string, code int32, err error) *BusinessError {
	return &BusinessError{
		LogId:       logId,
		Code:        code,
		OriginError: err,
	}
}

func (e *BusinessError) Error() string {
	if e == nil {
		return ""
	}
	return fmt.Sprintf("%s - %s", e.LogId, errcode.GetErrMsg(e.Code))
}

package errcode

import "fmt"

const (
	FIZZY_OCR_PARAM_VERIFY_FAILED = 1100 //业务参数校验不通过
)

var FizzyOcrErrorMsg = map[int32]string{
	FIZZY_OCR_PARAM_VERIFY_FAILED: "image_url和image_base64必须有一个不为空",
}

func FizzyOcrErrCodeInit() {
	for key, value := range FizzyOcrErrorMsg {
		if _, ok := errorMsg[key]; ok {
			panic(fmt.Sprintf("fizzy-ocr %d 错误码重复定义", key))
		}
		errorMsg[key] = value
	}
}

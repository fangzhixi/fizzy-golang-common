package errcode

import "fmt"

const (
	FIZZY_OCR_PARAM_VERIFY_FAILED_        = 1100 //业务参数校验不通过
	FIZZY_OCR_IMAGE_URL_PARAM_NOT_NULL    = 1101 //image_url不能为空
	FIZZY_OCR_IMAGE_BASE64_PARAM_NOT_NULL = 1102 //image_base64不能为空
)

var FizzyOcrErrorMsg = map[int32]string{
	FIZZY_OCR_PARAM_VERIFY_FAILED_:        "业务参数校验不通过",
	FIZZY_OCR_IMAGE_URL_PARAM_NOT_NULL:    "image_url不能为空",
	FIZZY_OCR_IMAGE_BASE64_PARAM_NOT_NULL: "image_base64不能为空",
}

func FizzyOcrErrCodeInit() {
	for key, value := range FizzyOcrErrorMsg {
		if _, ok := ErrorMsg[key]; ok {
			panic(fmt.Sprintf("fizzy-ocr %d 错误码重复定义", key))
		}
		ErrorMsg[key] = value
	}
}

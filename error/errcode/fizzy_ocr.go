package errcode

const (
	/******************** 参数检查错误码 (1100~1999) ********************/

	FIZZY_OCR_PARAM_VERIFY_IMAGE_FAILED int32 = 1101 //image_url和image_base64必须有一个不为空

	/******************** 业务逻辑错误码 (2100~2999) ********************/
)

var fizzyOcrErrorMsg = map[int32]string{
	FIZZY_OCR_PARAM_VERIFY_IMAGE_FAILED: "image_url和image_base64必须有一个不为空",
}

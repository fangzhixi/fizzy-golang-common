package errcode

/*
 * @Author       : zhixi.fang
 * @Date         : 2022-05-03 15:48:04
 * @LastEditors  : zhixi.fang
 * @LastEditTime : 2022-06-11 15:04:50
 */

const (
	/******************** 参数检查错误码 (1100~1999) ********************/

	FIZZY_OCR_PARAM_VERIFY_IMAGE_FAILED      = 1101 //image_url和image_base64必须有一个不为空
	FIZZY_OCR_PARAM_BASE64_LENGTH_TOO_SHORT  = 1102 //base64长度过短
	FIZZY_OCR_PARAM_BASE64_INCORRECT_PADDING = 1103 //base64解析失败, 填充不正确
	FIZZY_OCR_PARAM_VERIFY_LANG_FAILED       = 1104 //不支持识别的语种

	/******************** 业务逻辑错误码 (2100~2999) ********************/
	FIZZY_OCR_PICTURE_DOWNLOAD_FAILED            = 2101 //图片下载失败
	FIZZY_OCR_EXPECTED_CONTENT_NOT_DETECTED      = 2102 //未检测出预期内容
	FIZZY_OCR_UNABLE_TO_PARSE_BASE64_FILE_FORMAT = 2103 //参数无携带base64标准头, 无法解析文件格式
)

var fizzyOcrErrorMsg = map[int32]string{
	// 参数检查错误码 (1100~1999)

	FIZZY_OCR_PARAM_VERIFY_IMAGE_FAILED:      "image_url和image_base64必须有一个不为空",
	FIZZY_OCR_PARAM_BASE64_LENGTH_TOO_SHORT:  "base64长度过短",
	FIZZY_OCR_PARAM_BASE64_INCORRECT_PADDING: "base64解析失败, 填充不正确",
	FIZZY_OCR_PARAM_VERIFY_LANG_FAILED:       "不支持识别的语种",

	//业务逻辑错误码 (2100~2999)

	FIZZY_OCR_PICTURE_DOWNLOAD_FAILED:            "图片下载失败",
	FIZZY_OCR_EXPECTED_CONTENT_NOT_DETECTED:      "未检测出预期内容",
	FIZZY_OCR_UNABLE_TO_PARSE_BASE64_FILE_FORMAT: "参数无携带base64标准头, 无法解析文件格式",
}

package errcode

/*
 * @Author       : zhixi.fang
 * @Date         : 2022-06-11 10:08:42
 * @LastEditors  : zhixi.fang
 * @LastEditTime : 2022-06-11 15:04:58
 */

import (
	"fmt"

	"github.com/fangzhixi/go-common/define"
)

//错误码初始化
func Init(appName string) {
	var appErrorMsg map[int32]string

	switch appName {
	case define.OCR_API:
		appErrorMsg = fizzyOcrErrorMsg
	default:
		panic(fmt.Sprintf("app_name未匹配: %s", appName))
	}

	for key, value := range appErrorMsg {
		if _, ok := errorMsg[key]; ok {
			panic(fmt.Sprintf("%s 错误码重复定义: %d", appName, key))
		}
		fizzyOcrErrorMsg[key] = value
	}
}

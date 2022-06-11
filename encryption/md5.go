package encryption

/*
 * @Author       : zhixi.fang
 * @Date         : 2022-06-11 10:15:43
 * @LastEditors  : zhixi.fang
 * @LastEditTime : 2022-06-11 15:04:43
 */

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(message string) string {
	md5 := md5.New()
	md5.Write([]byte(message))
	return hex.EncodeToString(md5.Sum(nil))
}

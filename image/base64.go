package image

/*
 * @Author       : zhixi.fang
 * @Date         : 2022-06-11 10:08:42
 * @LastEditors  : zhixi.fang
 * @LastEditTime : 2022-06-11 15:05:18
 */

import (
	"strings"
)

func CutImgBase64Hander(imageBase64 string) string {
	if len(imageBase64) > 21 {
		if strings.Contains(imageBase64[:22], "data:image/jpg;base64,") { //JPG、JPEG
			return imageBase64[22:]
		} else if strings.Contains(imageBase64[:23], "data:image/jpeg;base64,") { //JPG、JPEG
			return imageBase64[23:]
		} else if strings.Contains(imageBase64[:22], "data:image/png;base64,") { //PNG
			return imageBase64[22:]
		} else if strings.Contains(imageBase64[:28], "data:application/pdf;base64,") { //PDF
			return imageBase64[28:]
		} else if strings.Contains(imageBase64[:37], "data:application/octet-stream;base64,") { //OFD
			return imageBase64[37:]
		}
	}
	return imageBase64
}

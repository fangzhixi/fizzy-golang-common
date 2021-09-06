package encryption

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(message string) string {
	md5 := md5.New()
	md5.Write([]byte(message))
	return hex.EncodeToString(md5.Sum(nil))
}

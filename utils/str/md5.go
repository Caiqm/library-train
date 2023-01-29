package str

import (
	"crypto/md5"
	"encoding/hex"
)

// md5加密
func StringToMd5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

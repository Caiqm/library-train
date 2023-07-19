package str

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"os"
)

type Encrypt struct {
}

// md5加密
func (e Encrypt) StringToMd5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// sha1加密
func (e Encrypt) StringToSha1(str string) string {
	o := sha1.New()
	o.Write([]byte(str))
	return hex.EncodeToString(o.Sum(nil))
}

// 计算文件sha1
func (e Encrypt) Sha1File(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	hash := sha1.New()
	hash.Write(data)
	return hex.EncodeToString(hash.Sum(nil)), nil
}

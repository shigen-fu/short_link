package utils

import (
	"crypto/md5"
	"encoding/hex"
)

const (
	length = 6
)

func GenerateCode(originURL string) string {
	// 计算MD5值
	h := md5.New()
	h.Write([]byte(originURL))
	md5Sum := hex.EncodeToString(h.Sum(nil))

	// 返回MD5值作为唯一的code
	return md5Sum[:length]
}

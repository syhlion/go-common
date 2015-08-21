package common

import (
	"crypto/md5"
	"encoding/hex"
)

func EncodeMd5(str string) string {
	m := md5.New()
	m.Write([]byte(str))
	return hex.EncodeToString(m.Sum(nil))
}

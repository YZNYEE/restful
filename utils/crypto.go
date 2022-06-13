package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func Cryptomd5(input string) string {
	m := md5.New()
	m.Write([]byte(input))
	res := hex.EncodeToString(m.Sum(nil))
	return res
}

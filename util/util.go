package util

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"io"
)

// Hash 用于计算给定字符串的哈希值的整数形式。
// 本函数实现了BKDR哈希算法。
func Hash(str string) uint64 {
	seed := uint64(13131)
	var hash uint64
	for i := 0; i < len(str); i++ {
		hash = hash*seed + uint64(str[i])
	}
	return (hash & 0x7FFFFFFFFFFFFFFF)
}

// GUID 获取GUID
func GUID() string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}

	h := md5.New()
	if _, err := h.Write([]byte(base64.URLEncoding.EncodeToString(b))); err != nil {
		return ""
	}
	return hex.EncodeToString(h.Sum(nil))
}

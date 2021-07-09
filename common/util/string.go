package util

import (
	"crypto/md5"
	"math/rand"
	"time"
)

func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return string(h.Sum(nil))
}

func RandString(len int) string {
	bytes := make([]byte, len)
	r := rand.New(rand.NewSource(time.Now().Unix()))
    for i := 0; i < len; i++ {
        b := r.Intn(26) + 65
        bytes[i] = byte(b)
    }
    return string(bytes)
}

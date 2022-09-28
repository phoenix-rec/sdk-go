package phoenix

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"strings"
)

// RandAllString 获取随机字符串
func RandAllString(lenNum int) string {
	var CHARS = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
		"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	str := strings.Builder{}
	length := len(CHARS)
	for i := 0; i < lenNum; i++ {
		l := CHARS[rand.Intn(length)]
		str.WriteString(l)
	}
	return str.String()
}

func CalSignature(customerId, timeNow, accessId, accessKey, nonce string, httpBody []byte) string {
	shaHash := sha256.New()
	shaHash.Write([]byte(customerId))
	shaHash.Write([]byte(accessId))
	shaHash.Write([]byte(accessKey))
	shaHash.Write([]byte(timeNow))
	shaHash.Write([]byte(nonce))
	shaHash.Write(httpBody)
	//将sha256转成16进制字符串
	signature := fmt.Sprintf("%x", shaHash.Sum(nil))
	return signature
}

package trait

import (
	"math/rand"
	"time"
)

/*RandomKind RandomKind */
type RandomKind int

/*random kinds */
const (
	RandomNum      RandomKind = iota // 纯数字
	RandomLower                      // 小写字母
	RandomUpper                      // 大写字母
	RandomLowerNum                   // 数字、小写字母
	RandomUpperNum                   // 数字、大写字母
	RandomAll                        // 数字、大小写字母
)

/*RandomString defines */
var (
	RandomString = map[RandomKind]string{
		RandomNum:      "0123456789",
		RandomLower:    "abcdefghijklmnopqrstuvwxyz",
		RandomUpper:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		RandomLowerNum: "0123456789abcdefghijklmnopqrstuvwxyz",
		RandomUpperNum: "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		RandomAll:      "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
	}
)

//GenerateRandomString 随机字符串
func GenerateRandomString(size int, kind ...RandomKind) string {
	s := RandomString[RandomAll]
	if kind != nil {
		if k, b := RandomString[kind[0]]; b == true {
			s = k
		}
	}
	var r []byte
	rand.Seed(time.Now().UnixNano())
	ls := len(s)
	for i := 0; i < size; i++ {
		r = append(r, s[rand.Intn(ls)])
	}
	return string(r)
}

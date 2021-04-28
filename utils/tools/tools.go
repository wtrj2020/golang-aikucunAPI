package tools

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/unknwon/com"
	"math/rand"
	"strconv"
	"time"
)

func Md5_salt(str, salt string) string {
	h := md5.New()
	h.Write([]byte(str + salt))
	md5res := hex.EncodeToString(h.Sum(nil))
	return md5res
}

func Guid(userId int) string {
	return Md5_salt(com.StrTo(userId).String(), strconv.FormatInt(time.Now().UnixNano(), 10))
}

func Rand(i int) int {
	rad := rand.New(rand.NewSource(time.Now().Unix()))
	return rad.Intn(i)
}

func GoodsIdRepeat(goodsIDs []int) []int {
	var secondInt []int
	for _, value := range goodsIDs {
		secondInt = append(secondInt, value)
	}
	fmt.Println(secondInt)
	return secondInt
}

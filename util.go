package wxpay

import (
	"math/rand"
	"strconv"
	"time"
)

const (
	KC_RAND_KIND_NUM   = iota // 纯数字
	KC_RAND_KIND_LOWER        // 小写字母
	KC_RAND_KIND_UPPER        // 大写字母
	KC_RAND_KIND_ALL          // 数字、大小写字母
)

func RandomNum(min int64, max int64) int64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	num := min + r.Int63n(max-min+1)
	return num
}

func RandomNumString(min int64, max int64) string {
	num := RandomNum(min, max)
	return strconv.FormatInt(num, 10)
}

func randStr(size int, kind int) string {
	ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	isAll := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if isAll {
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return string(result)
}

func GenRandStr(size int) string {
	if size <= 0 {
		return ""
	}
	return randStr(size, KC_RAND_KIND_ALL)
}

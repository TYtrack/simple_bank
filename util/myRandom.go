/*
 * @Author: your name
 * @Date: 2021-12-14 18:51:19
 * @LastEditTime: 2021-12-14 19:04:49
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /goproject/src/go_code/银行项目/util/Random.go
 */

package util

import (
	"math/rand"
	"strings"
	"time"
)

var alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt64(min int64, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var builder strings.Builder
	for i := 0; i < n; i++ {
		builder.WriteByte(alphabet[rand.Intn(len(alphabet))])

	}
	return builder.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomBalance() int64 {
	return RandomInt64(0, 100000)
}

func RandomCurrency() string {
	currencyList := []string{"EUR", "USD", "CNY"}
	return currencyList[rand.Intn(len(currencyList))]
}

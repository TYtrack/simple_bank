/*
 * @Author: your name
 * @Date: 2021-12-14 18:51:19
 * @LastEditTime: 2021-12-29 13:39:17
 * @LastEditors: TYtrack
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /bank_project/util/myRandom.go
 */

package util

import (
	"fmt"
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

func RandomEmail() string {

	return fmt.Sprintf("%v@email.com", RandomString(6))
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

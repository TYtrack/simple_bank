/*
 * @Author: TYtrack
 * @Description: ...
 * @Date: 2021-12-28 16:09:52
 * @LastEditors: TYtrack
 * @LastEditTime: 2021-12-28 16:11:05
 * @FilePath: /bank_project/util/currency.go
 */

package util

const (
	USD = "USD"
	EUR = "EUR"
	CNY = "CNY"
)

func IsSupportCurrency(currency string) bool {
	switch currency {
	case USD, EUR, CNY:
		return true
	}
	return false

}

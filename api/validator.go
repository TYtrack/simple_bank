/*
 * @Author: TYtrack
 * @Description: ...
 * @Date: 2021-12-28 16:06:18
 * @LastEditors: TYtrack
 * @LastEditTime: 2021-12-28 16:15:30
 * @FilePath: /bank_project/api/validator.go
 */

package api

import (
	"bank_project/util"

	"github.com/go-playground/validator/v10"
)

// 自定义一个验证器，并在serve.go中注册
var validCurrency validator.Func = func(fl validator.FieldLevel) bool {
	if currency, ok := fl.Field().Interface().(string); ok {
		// validate the currency correct
		return util.IsSupportCurrency(currency)
	}
	return false
}

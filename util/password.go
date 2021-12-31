/*
 * @Author: TYtrack
 * @Description: ...
 * @Date: 2021-12-29 19:18:39
 * @LastEditors: TYtrack
 * @LastEditTime: 2021-12-29 19:20:55
 * @FilePath: /bank_project/util/password.go
 */

package util

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPwd), err
}

func CheckPassword(hashedPwd string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(password))
}

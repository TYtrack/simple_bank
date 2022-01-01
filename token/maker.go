/*
 * @Author: TYtrack
 * @Description: ...
 * @Date: 2021-12-31 00:18:21
 * @LastEditors: TYtrack
 * @LastEditTime: 2021-12-31 11:31:49
 * @FilePath: /bank_project/token/maker.go
 */

package token

import "time"

type Maker interface {
	CreateToken(username string, duration time.Duration) (token string, err error)
	VerifyToken(token string) (payload *Payload, err error)
}

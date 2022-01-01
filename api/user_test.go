/*
 * @Author: TYtrack
 * @Description: ...
 * @Date: 2021-12-29 17:03:26
 * @LastEditors: TYtrack
 * @LastEditTime: 2022-01-01 16:03:09
 * @FilePath: /bank_project/api/user_test.go
 */

package api

import (
	mockdb "bank_project/db/mock"
	db "bank_project/db/sqlc"
	"bank_project/util"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

type eqCreateUserParamsMatcher struct {
	arg      db.CreateUserParams
	password string
}

func (e eqCreateUserParamsMatcher) Matches(x interface{}) bool {
	arg, ok := x.(db.CreateUserParams)
	if !ok {
		return false
	}
	err := util.CheckPassword(arg.HashPassword, e.password)
	if err != nil {
		return false
	}

	e.arg.HashPassword = arg.HashPassword
	return reflect.DeepEqual(e.arg, arg)
}

func (e eqCreateUserParamsMatcher) String() string {
	return fmt.Sprintf("matches arg %v and password %v", e.arg, e.password)

}

func EqCreateUserParams(arg db.CreateUserParams, pwd string) gomock.Matcher {
	return eqCreateUserParamsMatcher{arg, pwd}
}

func TestCreateUserApi(t *testing.T) {
	user, password := randomUser(t)

	testCases := []struct {
		name string
		body gin.H
		//使用这个函数来构建存根
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name: "ok",
			body: gin.H{
				"username": user.Username,
				"email":    user.Email,
				"fullname": user.FullName,
				"password": password,
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.CreateUserParams{
					Username: user.Username,
					Email:    user.Email,
					FullName: user.FullName,
				}
				store.EXPECT().
					CreateUser(gomock.Any(), EqCreateUserParams(arg, password)).
					Times(1).
					Return(user, nil)

			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recoder.Code)

				requireBodyMatchUser(t, recoder.Body, user)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			//这个finish会检查所有预期被调用的代码是否都被调用
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)

			tc.buildStubs(store)

			server := newTestServer(t, store)
			//使用httptest来记录http请求的响应
			recoder := httptest.NewRecorder()
			url := "/user"

			byte, err := json.Marshal(tc.body)
			require.NoError(t, err)

			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(byte))
			require.NoError(t, err)

			//通过路由器发送API请求，并将响应记录在recoder中
			server.router.ServeHTTP(recoder, request)

			tc.checkResponse(t, recoder)

		})

	}

}

func requireBodyMatchUser(t *testing.T, body *bytes.Buffer, user db.User) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var gotUser db.User
	err = json.Unmarshal(data, &gotUser)

	require.NoError(t, err)
	require.Equal(t, user.Username, gotUser.Username)
	require.Equal(t, user.FullName, gotUser.FullName)
	require.Equal(t, user.Email, gotUser.Email)
	require.Empty(t, gotUser.HashPassword)

}

func randomUser(t *testing.T) (db.User, string) {

	password := util.RandomString(6)
	hashed_password, err := util.HashPassword(password)
	require.NoError(t, err)

	user := db.User{
		Username:     util.RandomString(6),
		FullName:     util.RandomString(6),
		Email:        util.RandomEmail(),
		HashPassword: hashed_password,
	}
	return user, hashed_password

}

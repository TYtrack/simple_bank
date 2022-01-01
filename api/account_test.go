/*
 * @Author: TYtrack
 * @Description: ...
 * @Date: 2021-12-27 23:03:39
 * @LastEditors: TYtrack
 * @LastEditTime: 2022-01-01 20:45:13
 * @FilePath: /bank_project/api/account_test.go
 */

package api

import (
	mockdb "bank_project/db/mock"
	db "bank_project/db/sqlc"
	"bank_project/token"
	"bank_project/util"
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

// 检测模拟发送的请求和服务器接受的请求是否一样
func TestCreateAccountApi(t *testing.T) {
	user, _ := randomUser(t)
	account := randomAccount(user.Username)

	testCases := []struct {
		name      string
		accountID int64
		setupAuth func(t *testing.T, request *http.Request, tokenMaker token.Maker)

		//使用这个函数来构建存根
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name:      "ok",
			accountID: account.ID,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				//希望GetAccountForUpdate函数调用一次
				store.EXPECT().
					GetAccountForUpdate(gomock.Any(), gomock.Eq(account.ID)).
					Times(1).
					Return(account, nil)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recoder.Code)

				requireBodyMatchAccount(t, recoder.Body, account)
			},
		}, {
			name:      "NotFound",
			accountID: account.ID,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				//希望GetAccountForUpdate函数调用一次
				store.EXPECT().
					GetAccountForUpdate(gomock.Any(), gomock.Eq(account.ID)).
					Times(1).
					Return(db.Account{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recoder.Code)
				// requireBodyMatchAccount(t, recoder.Body, account)
			},
		}, {
			name:      "InternalError",
			accountID: account.ID,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				//希望GetAccountForUpdate函数调用一次
				store.EXPECT().
					GetAccountForUpdate(gomock.Any(), gomock.Eq(account.ID)).
					Times(1).
					Return(account, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recoder.Code)
			},
		}, {
			//参数有问题
			name:      "BadRequest",
			accountID: 0,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				//希望GetAccountForUpdate函数调用一次
				store.EXPECT().
					GetAccountForUpdate(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recoder.Code)
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
			url := fmt.Sprintf("/account/%d", tc.accountID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)
			tc.setupAuth(t, request, server.tokenMaker)

			//通过路由器发送API请求，并将响应记录在recoder中
			server.router.ServeHTTP(recoder, request)
			tc.checkResponse(t, recoder)

		})

	}

}

func requireBodyMatchAccount(t *testing.T, body *bytes.Buffer, account db.Account) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var gotAccount db.Account
	err = json.Unmarshal(data, &gotAccount)
	require.NoError(t, err)

	require.Equal(t, gotAccount, account)
}

func randomAccount(owner string) db.Account {
	return db.Account{
		ID:       util.RandomInt64(1, 1000),
		Owner:    owner,
		Currency: util.RandomCurrency(),
		Balance:  util.RandomBalance(),
	}
}

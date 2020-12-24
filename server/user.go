/**
 * File              : user.go
 * Author            : qiwen <yangdongyong@qianxin.com>
 * Date              : 24.12.2020
 * Last Modified Date: 24.12.2020
 * Last Modified By  : qiwen <yangdongyong@qianxin.com>
 */
package server

import (
	"context"

	"github.com/qiwenilli/auv.kit.example/api/user"
)

type User struct {
	user.User
}

func (u *User) Reg(ctx context.Context, req *user.Reg_Request) (*user.Reg_Response, error) {

	return &user.Reg_Response{Msg: "xxxx123"}, nil
}

func (u *User) Login(ctx context.Context, req *user.Login_Request) (*user.Login_Response, error) {

	return &user.Login_Response{Msg: "xxxx456"}, nil
}

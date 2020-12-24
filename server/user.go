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

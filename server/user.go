package server

import (
	"context"

	"github.com/qiwenilli/auv.kit.example/api"
)

type User struct {
	api.User
}

func (u *User) Reg(ctx context.Context, req *api.Reg_Request) (*api.Reg_Response, error) {

	return &api.Reg_Response{Msg: "xxxx123"}, nil
}

func (u *User) Login(ctx context.Context, req *api.Login_Request) (*api.Login_Response, error) {

	return &api.Login_Response{Msg: "xxxx456"}, nil
}

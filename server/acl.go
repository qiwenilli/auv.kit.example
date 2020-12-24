package server

import (
	"context"

	"github.com/qiwenilli/auv.kit.example/api/acl"
)

type Acl struct {
	acl.Acl
}

func (a *Acl) List(ctx context.Context, req *acl.List_Request) (*acl.List_Response, error) {

	return &acl.List_Response{AssetsList: []string{"xxxx123", "123"}}, nil
}

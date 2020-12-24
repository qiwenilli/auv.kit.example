/**
 * File              : acl.go
 * Author            : qiwen <yangdongyong@qianxin.com>
 * Date              : 24.12.2020
 * Last Modified Date: 24.12.2020
 * Last Modified By  : qiwen <yangdongyong@qianxin.com>
 */
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

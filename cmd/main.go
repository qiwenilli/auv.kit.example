package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	aclapi "github.com/qiwenilli/auv.kit.example/api/acl"
	userapi "github.com/qiwenilli/auv.kit.example/api/user"

	business "github.com/qiwenilli/auv.kit.example/server"

	// "github.com/qiwenilli/auv.kit"
	"github.com/qiwenilli/auv.kit/discovery"
	"github.com/qiwenilli/auv.kit/server"
)

func main() {

	// 通过twirp框架生成服务定义
	// api
	user := &business.User{}
	srv1 := userapi.NewUserServer(user)

	acl := &business.Acl{}
	srv2 := aclapi.NewAclServer(acl)

	var opts []server.Opt
	opts = append(opts, server.WithServiceName("auv.mobi"))
	// opts = append(opts, server.WithDieHookFunc(dieHookFun))
	// opts = append(opts, server.WithMiddlewares(customeMiddleware))
	opts = append(opts, server.WithRatelimit(2))
	opts = append(opts, server.WithIpWhiteList([]string{"192.168.*.*"}))
	opts = append(opts, server.WithServices(srv1, srv2))

	// curl -X POST -H 'Content-Type: application/json' 'http://192.168.177.132:8080/twirp/user.api.User/Login' -d '{}'
	// curl -X POST -H 'Content-Type: application/json' 'http://192.168.177.132:8080/twirp/user.api.User/Reg' -d '{}'

	go func() {
		// 等待服务启动成功
		time.Sleep(time.Second * 2)

		// 客户端，通过服务发现获得ip
		for i := 0; i < 10; i++ {

			addr := discovery.GetServerWithName("auv.mobi")
			fmt.Println(addr)

			cli := userapi.NewUserJSONClient("http://"+addr, &http.Client{})
			//
			ctx := context.Background()
			resp, err := cli.Login(ctx, &userapi.Login_Request{Name: "qiwen", Phone: "130"})
			fmt.Println(resp, err)

		}
	}()

	srv := server.NewServer()
	srv.Run(opts...)
}

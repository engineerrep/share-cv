package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"sharecvapi/app/api/internal/config"
	"sharecvapi/app/api/internal/handler"
	"sharecvapi/app/api/internal/svc"
	"sharecvapi/common/errorx"
	"sharecvapi/common/mmdb"
	"sharecvapi/common/response"
)

var configFile = flag.String("f", "etc/core-api.yaml", "the config.go file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	logx.MustSetup(c.Log)

	server := rest.MustNewServer(c.RestConf, rest.WithCors(), rest.WithUnauthorizedCallback(response.TokenErrorResponse))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	// 自定义错误
	httpx.SetErrorHandler(func(err error) (int, interface{}) {
		switch e := err.(type) {
		case *errorx.CodeError:
			return http.StatusOK, e.Data()
		default:
			return http.StatusInternalServerError, nil
		}
	})

	mmdb.Init(c.App.IPDataPath)
	defer mmdb.Close()

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

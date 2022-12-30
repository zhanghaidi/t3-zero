package main

import (
	"flag"
	"fmt"
	"net/http"

	"t3-zero/front/internal/config"
	"t3-zero/front/internal/handler"
	"t3-zero/front/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/front-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf, rest.WithCustomCors(func(header http.Header) {
		header.Set("Access-Control-Allow-Headers", "Content-Type,User-Agent,Accept,Origin,token,device-type")
		header.Set("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,OPTIONS")
	}, func(w http.ResponseWriter) {
		header := w.Header()
		header.Set("Access-Control-Allow-Headers", "Content-Type,User-Agent,Accept,Origin,token,device-type")
		header.Set("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,OPTIONS")
	}))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

package main

import (
	//	"bytes"
	"github.com/buaazp/fasthttprouter"
	"github.com/zhongyuanzyh/config-center/controller"
	"log"
	//	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"
	"github.com/zhongyuanzyh/config-center/extend/conf"
)

var Logs = &conf.Levels{}

type handler fasthttp.RequestHandler

var routerPost = map[string]handler{
	"/apollo": apollo,
}

var routerGet = map[string]handler{
	"/nacos": controller.Nacos,
}

func apollo(ctx *fasthttp.RequestCtx) {

}

func main() {
	log.Print("开始监听请求....")
	router := fasthttprouter.New()

	for k, v := range routerPost {
		router.POST(k, fasthttp.RequestHandler(v))
	}

	for k, v := range routerGet {
		router.GET(k, fasthttp.RequestHandler(v))

	}

	fasthttp.ListenAndServe(":9988", router.Handler)
}

package router

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"github.com/zhongyuanzyh/config-center/controller"
	"log"
)

type handler fasthttp.RequestHandler

var routerPost = map[string]handler{
	"/apollo": apollo,
	"/nacos":  controller.NacosPost,
}

var routerGet = map[string]handler{
	"/nacos": controller.NacosGet,
}
var Router = fasthttprouter.New()

func apollo(ctx *fasthttp.RequestCtx) {

}
func SetUp() {
	log.Print("开始监听请求....")
	for k, v := range routerPost {
		Router.POST(k, fasthttp.RequestHandler(v))
	}
	for k, v := range routerGet {
		Router.GET(k, fasthttp.RequestHandler(v))
	}
}

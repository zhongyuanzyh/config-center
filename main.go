package main

import (
	"github.com/valyala/fasthttp"
	"github.com/zhongyuanzyh/config-center/router"
)

func main() {
	//路由初始化
	router.SetUp()
	fasthttp.ListenAndServe(":9988", router.Router.Handler)
}

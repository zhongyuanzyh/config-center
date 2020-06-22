package main

import (
	"github.com/valyala/fasthttp"
	"github.com/zhongyuanzyh/config-center/extend/db"
	"github.com/zhongyuanzyh/config-center/router"
)

func main() {
	//路由初始化
	router.SetUp()
	//mysql初始化
	db.SetUp()
	fasthttp.ListenAndServe(":9988", router.Router.Handler)
}

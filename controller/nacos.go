package controller

import (
	"github.com/valyala/fasthttp"
	"log"
)

var url = "http://nacos.ryanzhong.com"

func Nacos(ctx *fasthttp.RequestCtx) {
	dataId := ctx.QueryArgs().Peek("dataId")
	group := ctx.QueryArgs().Peek("group")
	urlPath := url + "/nacos/v1/cs/configs?dataId=" + string(dataId) + "&group=" + string(group)
	status, resp, err := fasthttp.Get(nil, urlPath)
	if err != nil {
		log.Println("请求失败：", err.Error())
		return
	}
	if status != fasthttp.StatusOK {
		log.Println("请求没有成功：", status)
	}
	ctx.Write(resp)
}

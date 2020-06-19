package controller

import (
	"github.com/valyala/fasthttp"
	"log"
)

var url = "http://nacos.ryanzhong.com"
var content string

// NacosGet Get方式查看配置
func NacosGet(ctx *fasthttp.RequestCtx) {
	dataId := ctx.QueryArgs().Peek("dataId")
	group := ctx.QueryArgs().Peek("group")
	urlPath := url + "/nacos/v1/cs/configs?dataId=" + string(dataId) + "&group=" + string(group)
	status, resp, err := fasthttp.Get(nil, urlPath)
	content = string(resp)
	if err != nil {
		log.Println("请求失败：", err.Error())
		return
	}
	if status != fasthttp.StatusOK {
		log.Println("请求没有成功：", status)
	}
	ctx.Write(resp)
	log.Println(content)
}

// NacosPost Post方式修改配置
func NacosPost(ctx *fasthttp.RequestCtx) {
	args := &fasthttp.Args{}
	args.Add("dataId", "nacos_golang")
	args.Add("group", "DEFAULT_GROUP")
	args.Add("content", content)
	urlPath := url + "/nacos/v1/cs/configs?"
	log.Println(content)
	status, resp, err := fasthttp.Post(nil, urlPath, args)

	if err != nil {
		log.Println("请求失败：", err.Error())
		return
	}
	if status != fasthttp.StatusOK {
		log.Println("请求没有成功：", status)
	}
	ctx.Write(resp)
}

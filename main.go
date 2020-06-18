package main

import (
	//	"bytes"
	"encoding/json"
	"fmt"
	"log"

	"github.com/buaazp/fasthttprouter"
	"github.com/ghodss/yaml"
	//	"github.com/spf13/viper"
	"extend/conf"
	"github.com/valyala/fasthttp"
)

var Logs = &conf.Levels{}

var url = "http://nacos.ryanzhong.com"

type handler fasthttp.RequestHandler

var routerPost = map[string]handler{
	"/apollo": apollo,
}

var routerGet = map[string]handler{
	"/nacos": nacos,
}

func nacos(ctx *fasthttp.RequestCtx) {
	dataId := ctx.QueryArgs().Peek("dataId")
	group := ctx.QueryArgs().Peek("group")
	urlPath := url + "/nacos/v1/cs/configs?dataId=" + string(dataId) + "&group=" + string(group)
	status, resp, err := fasthttp.Get(nil, urlPath)
	if err != nil {
		fmt.Println("请求失败：", err.Error())
		return
	}
	if status != fasthttp.StatusOK {
		fmt.Println("请求没有成功：", status)
	}
	//	viper.SetConfigType("YAML")
	//	err = viper.ReadConfig(bytes.NewBuffer(resp))
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	viper.UnmarshalKey("logging", Logs)
	j, err := yaml.YAMLToJSON(resp)
	if err != nil {
		log.Println("yamltojson错误", err.Error())
	}
	//fmt.Printf("%v", string(j))
	l := &Log{}

	json.Unmarshal(j, l)
	l.Logging.Level.Io.Swagger.Models.Parameters.AbstractSerializableParameter = "info"

	//fmt.Printf("%+v", *l)
	jb, _ := json.Marshal(l)
	//fmt.Println(string(jb))
	yb, _ := yaml.JSONToYAML(jb)

	ctx.Write(yb)
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

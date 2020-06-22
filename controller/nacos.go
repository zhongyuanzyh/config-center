package controller

import (
	"encoding/json"
	//"fmt"
	"log"

	//"github.com/ghodss/yaml"
	"github.com/valyala/fasthttp"
	"github.com/zhongyuanzyh/config-center/extend/db"
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
	//var respBody interface{}
	//t, _ := yaml.YAMLToJSON(resp)
	//json.Unmarshal(t, &respBody)
	//	log.Println(respBody)
	//fmt.Printf("%+v", respBody)

	ctx.Write(resp)
	//log.Println(content)
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

// NacosGetConfigNameAll获取nacos中所有的配置名
func NacosGetConigNameAll(ctx *fasthttp.RequestCtx) {
	rows, err := db.MySqlPool.Query("select data_id from config_info")
	if err != nil {
		log.Println("nacos配置数据库查询出错", err)
	}
	var count int = 1
	var configs map[int]string
	configs = make(map[int]string)
	for rows.Next() {
		var configName string
		rows.Scan(&configName)
		configs[count] = configName
		count++
		log.Printf("nacos配置名是：%s", configName)
	}
	for k, v := range configs {
		log.Println(k, v)
	}
	resp, _ := json.Marshal(configs)
	ctx.Write(resp)
}

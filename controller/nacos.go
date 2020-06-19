package controller

import (
	"github.com/ghodss/yaml"
	//	"github.com/spf13/viper"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"github.com/zhongyuanzyh/config-center/extend/conf"
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
	l := &conf.Log{}

	json.Unmarshal(j, l)
	l.Logging.Level.Io.Swagger.Models.Parameters.AbstractSerializableParameter = "info"

	//fmt.Printf("%+v", *l)
	jb, _ := json.Marshal(l)
	//fmt.Println(string(jb))
	yb, _ := yaml.JSONToYAML(jb)

	ctx.Write(yb)
}

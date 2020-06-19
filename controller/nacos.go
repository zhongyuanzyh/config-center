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
	logConf := &conf.Log{}
	mybatisConf := &conf.MyBatis{}

	json.Unmarshal(j, logConf)
	json.Unmarshal(j, mybatisConf)
	//修改应该是post，而不是在get里面这个和nacos那边的api要保持一致
	//l.Logging.Level.Io.Swagger.Models.Parameters.AbstractSerializableParameter = "info"

	//fmt.Printf("%+v", *l)
	jByteLog, _ := json.Marshal(logConf)
	jByteBatis, _ := json.Marshal(mybatisConf)
	//fmt.Println(string(jb))
	yByteLog, _ := yaml.JSONToYAML(jByteLog)
	yByteBatis, _ := yaml.JSONToYAML(jByteBatis)

	ctx.Write(yByteLog)
	ctx.Write(yByteBatis)
}

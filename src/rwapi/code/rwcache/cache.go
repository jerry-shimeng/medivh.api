package rwcache

import (
	"rwapi/models"
	"strings"
	"encoding/json"
)

var dataCache map[string]*[]models.AppDataModel

func init() {
	//启动缓存
	dataCache = make(map[string]*[]models.AppDataModel, 100)
}
//添加字符串消息
func AddMsg(msg string) {
	//判断非心跳数据
	msg = strings.Trim(msg, "\n")
	if len(msg) > 5 {
		//json转换
		var v []models.AppDataModel
		json.Unmarshal([]byte(msg), &v)
		if v != nil {
			add(v)
		}
	}
}
//添加对象模型消息
func add(v []models.AppDataModel) {
	if v == nil || len(v) == 0 {
		return
	}
	//循环获取
	for _, a := range v {
		if a.AppKey != "" {
			var key = a.AppKey
			//寻找list
			var list = fineKey(key)
			*list = append(*list, a)
		}
	}
}

func fineKey(key string) *[]models.AppDataModel {
	v, r := dataCache[key]
	if r && v != nil {
		return v
	}else {
		a := make([]models.AppDataModel, 0)
		v = &a
		dataCache[key] = v
	}
	return v
}

func GetDatas(appkey string) []models.AppDataModel {

	v, r := dataCache[appkey]

	if r  && v != nil {
		dataCache[appkey] = nil
		return *v
	}
	return nil
}

func GetAllKey() []string {
	var list = make([]string,0)

	for key, _:= range dataCache{
		list = append(list,key)
	}
	return list
}




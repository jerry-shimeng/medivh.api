package rwconnection

var dict map[string]string

func init() {
	dict = make(map[string]string,10)
}

//添加一个配置
func SetConnDict(id, appKey string)  {
	dict[appKey] = id;
}

func GetConnDict(appkey string)string  {
	var r = dict[appkey];
	return r
}

func RemoveConnDict(appkey string)  {
	delete(dict,appkey)
}

//判断是否存在
func ExitConnDict(appkey string)bool  {
	_ ,b := dict[appkey]
	return b
}
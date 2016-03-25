package rwconnection


var dict map[string][]string

func init() {
	dict = make(map[string][]string, 100)
}

//添加一个配置
func SetConnDict(id, appKey string) {
	//	dict[appKey] = id;
	if dict[appKey] == nil {
		dict[appKey] = make([]string, 0)
		dict[appKey] = append(dict[appKey], id)
	}else {
		dict[appKey] = append(dict[appKey], id)
	}
}


func RemoveConnDict(appkey, cid string) {
	//先获取列表
	var tmp = dict[appkey]
	if tmp != nil {
		var index = -1;
		//先找到索引
		for i, v := range tmp {
			if v == cid {
				index = i
				break
			}
		}
		if index > 0 {
			dict[appkey] = append(tmp[:index], tmp[index + 1:]...)
		}
	}
}

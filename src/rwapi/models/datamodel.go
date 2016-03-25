package models


//应用服务器数据模型
type AppDataModel struct {
	AppKey      string      //应用唯一标示 public key
	ModuleType  int         //模块类型 1一次性数据     3心跳数据
	CounterType int         //一次性数据类型  1err  3business 4custom
	Count       int         //数据总量
	Level       int         //数据级别
	Mark        string      //数据唯一标示
	CreateTime  int64       //创建时间
	Result      interface{} //心跳运行结果
}

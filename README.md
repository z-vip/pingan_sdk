# pingan_sdk
平安银行见证宝API

#调用方法
app,err := pingan_sdk.NewApp(config,redisPool)<br>
config:配置信息 map或 pingan_sdk.BaseApp{}<br>
redisPool:redis的连接池，如果不传默认为127.0.0.1:6379/0

#api调用，具体看文档和方法定义
app.func()

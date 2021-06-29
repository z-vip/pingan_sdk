package pingan_sdk

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/z-vip/pingan_sdk/util"
	"reflect"
	"time"
)

/**
生成新的应用
config定义app的配置
*/
func NewApp(config interface{}, args ...interface{}) (*App, error) {
	var appBase AppBase
	_ = util.Convert(config, &appBase)
	var app = &App{
		AppBase: appBase,
	}
	//检查配置参数是否未定义
	if err := app.AppInit(); err != nil {
		return app, err
	}
	//定义redis"
	if len(args) > 0 && reflect.TypeOf(args[0]).String() == "*redis.Pool" {
		app.Redis = args[0].(*redis.Pool)
		fmt.Println("==自定义redis==", app.Redis)
	} else {
		//默认本机redis
		app.Redis = &redis.Pool{
			MaxIdle:     50,
			MaxActive:   5000,
			IdleTimeout: 60 * time.Second,
			Dial: func() (redis.Conn, error) {
				c, err := redis.Dial("tcp", "127.0.0.1:6379")
				if err != nil {
					return nil, err
				}
				_, err = c.Do("SELECT", 0)
				if err != nil {
					_ = c.Close()
					return nil, err
				}
				return c, nil
			},
		}
	}
	return app, nil
}

package util

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// post 网络请求 ,params 是url.Values类型
func HttpPost(uri string, params url.Values) ([]byte, error) {
	response, err := http.PostForm(uri, params)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}

func HttpGet(uri string) ([]byte, error) {
	response, err := http.Get(uri)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}

func HttpRequest(method, url, data string, header map[string]string) (body []byte, err error) {
	//生成client 参数为默认
	client := &http.Client{}
	//提交请求
	var request *http.Request
	request, err = http.NewRequest(method, url, strings.NewReader(data))
	if err != nil {
		return
	}
	//增加header选项
	if len(header) > 0 {
		for k, v := range header {
			request.Header.Add(k, v)
		}
	}
	//处理返回结果
	response, err := client.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()
	body, err = ioutil.ReadAll(response.Body)
	return
}

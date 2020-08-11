package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

//	客户端
func main() {
	apiUrl := "http://127.0.0.1:9090/post"
	contentType := "application/json"
	data := `{"name": "小王子", "age": 18}`
	res, err := http.Post(apiUrl, contentType, strings.NewReader(data))
	if err != nil {
		fmt.Printf("get url failed, err: %v\n", err)
		return
	}
	defer res.Body.Close()
	//	从response中把服务端返回的数据读出来
	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("read response failed, err: %v\n", err)
		return
	}

}

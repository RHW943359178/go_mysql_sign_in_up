package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//	http server端

func formShow(w http.ResponseWriter, r *http.Request) {
	//	读取文件
	content, err := ioutil.ReadFile("./pages/form.html")
	if err != nil {
		//fmt.Printf("read file failed, err: %v\n", err)
		//	将错误类型转为字节输出到页面
		_, _ = w.Write([]byte(fmt.Sprintf("%v\n", err)))
		return
	}
	_, err = w.Write(content)
	if err != nil {
		fmt.Printf("send data failed, err: %v\n", err)
		return
	}
}

//	注册成功返回
func signUpSuccess() {

}

func formInput(w http.ResponseWriter, r *http.Request) {

}

func main() {
	fmt.Println("服务启动成功")
	http.HandleFunc("/sign_up/", formShow)
	http.HandleFunc("/sign_in/", formInput)
	_ = http.ListenAndServe("127.0.0.1:9090", nil)
}

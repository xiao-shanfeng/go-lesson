package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Body struct {
	Name string
}

func handle(res http.ResponseWriter, req *http.Request) {
	fmt.Println("请求方式:", req.Method)
	fmt.Println("请求地址:", req.URL.Path)
	//fmt.Println("请求参数:", req)
	// 在response中添加header信息
	header := res.Header()
	header["test"] = []string{"test", "test1"}
	fmt.Println(req.Header["Test"])
	switch req.Method {
	case "GET":
		fmt.Println("我是get请求")
	case "POST":
		fmt.Println("我是post请求")
		body := req.Body
		fmt.Println(body)
		// 通过ioutil.readAll获取POST请求体中的所有数据
		b, _ := ioutil.ReadAll(body)
		fmt.Println(b)
		fmt.Println(string(b))
		// 通过使用json把请求数据转换成结构体，用来获取请求数据
		bodyS := Body{
			Name: "",
		}
		_ = json.Unmarshal(b, &bodyS)
		fmt.Printf("%+v", bodyS)
		fmt.Println(bodyS.Name)
	}
	// 修改状态码
	res.WriteHeader(http.StatusBadRequest)
	res.Write([]byte("我请求成功了"))
}

func main() {
	mux := http.NewServeMux()
	//http.Handle()
	mux.HandleFunc("/test", handle)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println(err)
		return
	}
}

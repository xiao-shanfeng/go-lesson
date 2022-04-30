package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	client := &http.Client{}
	//req, _ := http.NewRequest("GET", "http://localhost:8080/test", nil)
	body := bytes.NewBuffer([]byte("{\"name\":\"Jack\"}"))
	req, _ := http.NewRequest("POST", "http://localhost:8080/test", body)
	head := req.Header
	head["test"] = []string{"test", "test1"}
	res, _ := client.Do(req)
	//fmt.Println(res)
	b, _ := ioutil.ReadAll(res.Body)
	//fmt.Println(b)
	fmt.Println(string(b))
}

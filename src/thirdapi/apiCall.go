package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main(){
	client := &http.Client{}
	reqest, _ := http.NewRequest("GET", "http://127.0.0.1:8001/ping", nil)
	reqest.Header.Set("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	reqest.Header.Set("Accept-Charset","GBK,utf-8;q=0.7,*;q=0.3")
	reqest.Header.Set("Accept-Encoding","gzip,deflate,sdch")
	reqest.Header.Set("Accept-Language","zh-CN,zh;q=0.8")
	reqest.Header.Set("Cache-Control","max-age=0")
	reqest.Header.Set("Connection","keep-alive")
	reqest.Header.Set("User-Agent","chrome 100")

	response,_ := client.Do(reqest)
	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		bodystr := string(body)
		fmt.Println(bodystr)
	}
}

/**
 * golang实现linux里面的curl
 * 目前实现了curl中 -d -H -e 方法，其他的后续慢慢支持
 * @author LiMao - limao777@126.com
 * @license MIT - https://opensource.org/licenses/MIT
 */
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {

	d := flag.String("d", "",
		"<data>   HTTP POST方式传送数据")
	H := flag.String("H",
		"", "<line>自定义头信息传递给服务器")
	e := flag.String("e",
		"", "来源网址")
	url := os.Args[len(os.Args)-1]

	//TODO 此处应对参数做合法性校验，以后再弄

	if !strings.Contains(strings.ToLower(url), "http") {
		url = "http://" + url
	}
	flag.Parse()

	/*
		fmt.Println("d:", *d)
		fmt.Println("H:", *H)
		fmt.Println("e:", *e)
		fmt.Println("os:", os.Args[len(os.Args)-1])
	*/
	client := &http.Client{}

	var req *http.Request
	var err error
	if *d == "" {
		req, err = http.NewRequest("GET", url, nil)
	} else {
		req, err = http.NewRequest("POST", url, strings.NewReader(*d))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded") //post方式需要设置Content-Type才能正确post数据
	}
	if err != nil {
		//TODO handle error
	}

	if *H != "" {
		//TODO 处理较为粗糙，以后改
		header_ext := strings.Split(*H, ":")
		req.Header.Set(header_ext[0], header_ext[1])
	}

	if *e != "" {
		req.Header.Set("Refer", *e)
	}
	//	req.Header.Set("Cookie", "name=anny")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//TODO handle error
	}

	fmt.Println(string(body))

}

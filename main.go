package main

import (
	"fmt"
	"jRebel-license-server/handler"
	"net/http"
)

func main() {
	//1.注册处理器函数
	http.HandleFunc("/uuid", handler.UUID)
	http.HandleFunc("/jrebel/leases", handler.Leases)
	http.HandleFunc("/jrebel/leases/1", handler.Leases1)
	http.HandleFunc("/jrebel/validate-connection", handler.ValidateConnection)
	fmt.Println(`
    启动成功 端口号:9000
	GET /uuid 生成随机串
	http://localhost:9000/{uuid} 放入jrebel激活地址栏`)
	err := http.ListenAndServe("127.0.0.1:9000", nil)
	if err != nil {
		fmt.Printf("http.ListenAndServe()函数执行错误,错误为:%v\n", err)
		return
	}

}

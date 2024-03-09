package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

/*
编写一个 HTTP 服务器
1.接受客户端 request，并将 request 中带的 header 写入 response header
2.读取当前系统的环境变量中的 VERSION 配置，并写入 response header
3.Server端记录访问日志包括客户端 ip，HTTP 返回码，输出到 server 端的标准输出
4.当访问 localhost/healthz 时，应返回 200
*/

func healthz(w http.ResponseWriter, r *http.Request) {
	// 将客户端的请求头复制到响应头
	for name, values := range r.Header {
		for _, value := range values {
			w.Header().Add(name, value)
		}
	}

	// 读取环境变量VERSION，并添加到响应头
	version := os.Getenv("VERSION")
	w.Header().Set("VERSION", version)

	// 记录访问日志
	fmt.Printf("Client IP: %s, HTTP Method: %s, Path: %s, Status Code: %d\n", r.RemoteAddr, r.Method, r.URL.Path, http.StatusOK)

	io.WriteString(w, "ok")
}

// HandlerFuncAdapter 适配器允许普通函数作为 HTTP 处理器
type HandlerFuncAdapter func(http.ResponseWriter, *http.Request)

// ServeHTTP 实现 http.Handler 接口
func (h HandlerFuncAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 调用函数本身
	h(w, r)
}

func main() {
	http.Handle("/healthz", HandlerFuncAdapter(healthz))
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	resp, err := http.Get("http://localhost/healthz")
	if err != nil {
		fmt.Println("Get Resp err:", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Get Resp body err:", err)
		return
	}
	fmt.Println("Resp body：" + string(body))
}

package main

import (
	"fmt"
	"net/http"
	"log"
	"os"
	"strings"
	"net"
)
func sayHello(w http.ResponseWriter, req *http.Request) {
	// 1.接收客户端 request，并将 request 中带的 header 写入 response header
	for key, values := range req.Header {
		for _, value := range values {
			w.Header().Set(key, value)
		}
	}
	// 2.读取当前系统的环境变量中的 VERSION 配置，并写入 response header
	version := os.Getenv("VERSION")
	w.Header().Set("version", version)

	// 3.Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
	clientIP := getClientIP(req)
	fmt.Printf("clientIP: %s\n", clientIP)
	fmt.Printf("http response code: %v\n", 200)
}

// 4.当访问 localhost/healthz 时，应返回 200
func healthz(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
}

func getClientIP(req *http.Request) string {
	xForwordedFor := req.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwordedFor, ",")[0])
	if ip != "" {
		return ip
	}
	ip = strings.TrimSpace(req.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(req.RemoteAddr)); err == nil {
		return ip
	}
	return ""
}

func main() {
	os.Setenv("VERSION", "v1.0")
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

package main

import (
    "fmt"
    "net/http"
    "strings"
    "log"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()  // 解析参数，默认是不会解析的
    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }
    fmt.Fprintf(w, "Hello World, I'm moemoefish!") // 这个写入到 w 的是输出到客户端的
}

func main() {
    http.HandleFunc("/", sayhelloName) // 设置访问的路由
    http.ListenAndServe(":8090", nil) // 设置监听的端口
}

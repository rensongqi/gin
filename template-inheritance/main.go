package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request)  {
	// 定义模板
	// 解析模板
	t, err := template.ParseFiles("./template/base.tmpl", "./template/index.tmpl")
	if err != nil {
		fmt.Printf("Parse templat failed. err=%#v\n", err)
		return
	}

	//自定义模板符号

	// 渲染模板
	name := "小王子"
	t.ExecuteTemplate(w, "index.tmpl", name)
}

func home(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	// 解析模板
	t, err := template.ParseFiles("./template/base.tmpl", "./template/home.tmpl")
	if err != nil {
		fmt.Printf("Parse templat failed. err=%v\n", err)
		return
	}
	// 渲染模板
	name := "RSQ"
	t.ExecuteTemplate(w, "home.tmpl", name)
}

func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/home", home)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("HTTP server start failed. err=%v\n", err)
		return
	}
}
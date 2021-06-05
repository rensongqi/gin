package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request)  {
	//定义函数
	k := func(name string) (string, error){
		return name + "真帅", nil
	}
	//2. 解析模板
	t := template.New("f.tmpl")
	// 告诉模板引擎，现在多了一个自定义函数
	t.Funcs(template.FuncMap{
		"kua": k,
	})
	t.ParseFiles("../template/f.tmpl")

	//3. 渲染模板
	name := "小王子"
	err := t.Execute(w, name)
	if err != nil {
		fmt.Printf("Render templat failed. err=%v\n", err)
		return
	}
}

func demo1(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	// 解析模板
	t, err := template.ParseFiles("../template/t.tmpl", "../template/ul.tmpl")
	if err != nil {
		fmt.Printf("Parse templat failed. err=%v\n", err)
		return
	}
	// 渲染模板
	name := "小王子"
	t.Execute(w, name)

}

func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/tmplDemo", demo1)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("HTTP server start failed. err=%v\n", err)
		return
	}
}
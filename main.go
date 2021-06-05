package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name string
	Age int
	Gender string
}

func sayHello(w http.ResponseWriter, r *http.Request)  {
	//2. 解析模板
	t, err := template.ParseFiles("template/hello.tmpl")
	if err != nil {
		fmt.Printf("Parse templat failed. err=%v\n", err)
		return
	}

	//3. 渲染模板
	u1 := User{
		Name: "小王子",
		Age: 18,
		Gender: "男",
	}
	m1 := map[string]interface{}{
		"name": "小王子",
		"age": 30,
		"gender": "男",
	}
	hobbyList := []string{
		"篮球",
		"足球",
		"双色球",
	}
	err = t.Execute(w, map[string]interface{}{
		"u1": u1,
		"m1": m1,
		"hobby": hobbyList,
	})
	if err != nil {
		fmt.Printf("Render templat failed. err=%v\n", err)
		return
	}

}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("HTTP server start failed. err=%v\n", err)
		return
	}
	//strings.Contains()
}
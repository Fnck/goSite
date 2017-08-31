package main

import (
	"fmt"
	"qiaobo.com/anotherProject"
	"io"
	"net/http"
)

func main()  {
	http.HandleFunc("/", work)
	http.ListenAndServe(":80", nil)
	fmt.Println("web server is start!")
}

func work(w http.ResponseWriter, r *http.Request)  {
	io.WriteString(w, anotherProject.Reverse("hello world"))
}


package main

import (
	"fmt"
	"net/http"
	"strings"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func func_temp(w http.ResponseWriter, r *http.Request) {
	res := strings.Builder{}
	res.WriteString(fmt.Sprintf("r.Host:\t%s \n", r.Host))
	res.WriteString(fmt.Sprintf("r.RemoteAddr:\t%s \n", r.RemoteAddr))
	res.WriteString(fmt.Sprintf("r.RequestURI:\t%s \n", r.RequestURI))
	if r.TLS != nil {
		res.WriteString(fmt.Sprintf("r.TLS.ServerName:\t%s \n", r.TLS.ServerName))
	}

	res.WriteString(fmt.Sprintf("r.URL.RawPath:\t%s \n", r.URL.RawPath))
	res.WriteString(fmt.Sprintf("r.URL.User.Username():\t%s \n", r.URL.User.Username()))
	res.WriteString(fmt.Sprintf("r.URL.User.String():\t%s \n", r.URL.User.String()))
	res.WriteString(fmt.Sprintf("r.URL.Path:\t%s \n", r.URL.Path))
	res.WriteString(fmt.Sprintf("r.Proto:\t%s \n", r.Proto))
	res.WriteString(fmt.Sprintf("r.Close:\t%t \n", r.Close))
	res.WriteString(fmt.Sprintf("r.Method:\t%s \n", r.Method))
	res.WriteString(fmt.Sprintf("r.ContentLength:\t%d \n", r.ContentLength))
	res.WriteString(fmt.Sprintf("r.TransferEncoding:\t%s \n", strings.Join(r.TransferEncoding, ",")))

	res.WriteString(fmt.Sprintln("\n"))
	res.WriteString(fmt.Sprintf("r.Header(头信息):\n"))
	for k, v := range r.Header {
		// k 是 string， v 是[]string ，经过测试，长度全是1
		res.WriteString(fmt.Sprintf("k: %s, v: %s\n", k, v))
	}

	res.WriteString(fmt.Sprintln("\n"))
	res.WriteString(fmt.Sprintf("r.Form(表单信息):\n"))
	for k, v := range r.Form {
		// k 是 string， v 是[]string ，经过测试，长度全是1
		res.WriteString(fmt.Sprintf("k: %s, v: %s\n", k, v))
	}

	res.WriteString(fmt.Sprintln("\n"))
	res.WriteString(fmt.Sprintf("r.PostForm(POST表单信息):\n"))
	for k, v := range r.PostForm {
		// k 是 string， v 是[]string ，经过测试，长度全是1
		res.WriteString(fmt.Sprintf("k: %s, v: %s\n", k, v))
	}

	res.WriteString(fmt.Sprintln("\n"))
	res.WriteString(fmt.Sprintf("r.PostForm(POST表单信息):\n"))
	for k, v := range r.PostForm {
		// k 是 string， v 是[]string ，经过测试，长度全是1
		res.WriteString(fmt.Sprintf("k: %s, v: %s\n", k, v))
	}

	/* 带文件的表单 */
	if r.MultipartForm != nil {
		res.WriteString(fmt.Sprintln("\n"))
		res.WriteString(fmt.Sprintf("r.MultipartForm.File(带文件的表单信息file):\n"))
		for k, v := range r.MultipartForm.File {
			// k 是 string， v 是[]string ，经过测试，长度全是1
			res.WriteString(fmt.Sprintf("k: %s, v: %s\n", k, v))
		}
		res.WriteString(fmt.Sprintln("\n"))
		res.WriteString(fmt.Sprintf("r.MultipartForm.Value(带文件的表单信息value):\n"))
		for k, v := range r.MultipartForm.Value {
			// k 是 string， v 是[]string ，经过测试，长度全是1
			res.WriteString(fmt.Sprintf("k: %s, v: %s\n", k, v))
		}
	}

	res.WriteString(fmt.Sprintln("\n"))
	res.WriteString(fmt.Sprintf("r.Trailer:\n"))
	for k, v := range r.Trailer {
		// k 是 string， v 是[]string ，经过测试，长度全是1
		res.WriteString(fmt.Sprintf("k: %s, v: %s\n", k, v))
	}

	if r.Response != nil {
		res.WriteString(fmt.Sprintln("\n"))
		res.WriteString(fmt.Sprintf("r.Response:\n"))
		res.WriteString(fmt.Sprintf("r.Response.Status:%s\n", r.Response.Status))
		res.WriteString(fmt.Sprintf("r.Response.Status:%d\n", r.Response.StatusCode))

	}

	fmt.Fprintln(w, res.String())
}

func a(a, b int) (c int){
	return a + b
}


type func_a func(a int, b int)(c int)

func func_c(a, b int, f func_a)(c int){
	return f(a, b)
}

func (f_a func_a)func_a_low(a, b int){
	f_a(a, b)
}

type greeting string

func (g greeting) SayHello(){
	fmt.Println("hello")
}

func (g greeting) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, g)
}

func main() {




	c := func_c(1,2, a)
	fmt.Println(c)

	// 实现了ServeHTTP方法的Handler接口
	http.Handle("/greeting", greeting("Welcome, dj"))


	http.HandleFunc("/tmp", func_temp)
	http.HandleFunc("/", index)

	/*
	http.Server{
		Addr:              "",
		Handler:           nil,
		TLSConfig:         nil,
		ReadTimeout:       0,
		ReadHeaderTimeout: 0,
		WriteTimeout:      0,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}

	*/
	server := http.Server{
		Addr: "0.0.0.0:8080",
	}
	server.ListenAndServe()
	//
	http.ListenAndServe("0.0.0.0:8080", nil)
}

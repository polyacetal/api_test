package main

import (
	"fmt"
	"net/http"
)

type HelloHandler struct{}
type HogeHandler struct{}
type FugaHandler struct{}

// *HelloHandler がインターフェース http.Handler を実装
func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

func (h *HogeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hoge")
}

func (h *FugaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "fuga")
}

func main() {
	// HelloHandler 型の変数を宣言
	hello := HelloHandler{}
	hoge := HogeHandler{}
	fuga := FugaHandler{}

	server := http.Server{
		Addr: ":8080",
		Handler: nil,	// DefaultServeMux を使用
	}

	// DefaultServeMux にハンドラを付与
	http.Handle("/hello", &hello)
	http.Handle("/hoge", &hoge)
	http.Handle("/fuga", &fuga)

	server.ListenAndServe()
}

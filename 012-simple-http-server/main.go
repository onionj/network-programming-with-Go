package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
)

func main() {
	filieServer := http.FileServer(http.Dir("./static/"))
	http.Handle("/", filieServer)

	http.HandleFunc("/hello/", hello)

	err := http.ListenAndServe(":8000", nil)
	checkError(err)
}

func hello(w http.ResponseWriter, r *http.Request) {
	dump, _ := httputil.DumpRequest(r, true)
	fmt.Println(string(dump))

	w.Write([]byte("Hello World"))
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}

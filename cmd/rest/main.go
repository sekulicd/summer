package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", hello)
	fmt.Println("server started ...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println(err)
	}
}


func hello(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello world seki")
	if err != nil {
		fmt.Println(err)
	}
}

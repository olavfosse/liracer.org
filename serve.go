package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("docs")))
	fmt.Println("http://localhost:3210")
	fmt.Print(http.ListenAndServe(":3210", nil))
}

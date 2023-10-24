package main

import (
	"fmt"
	"go-web-app7/pkg"
	"net/http"
)

const portNumber = ":8080"

// main is the main function
func main() {
	http.HandleFunc("/", pkg.Home)
	http.HandleFunc("/about", pkg.About)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}

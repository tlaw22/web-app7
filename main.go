package main

import (
	"fmt"
	"go-web-app7/pkg"
	"log"
	"net/http"
)

const portNumber = ":8080"

// main is the main function
func main() {

	var app pkg.AppConfig
	tc, err := pkg.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	pkg.NewTemplate(&app)

	http.HandleFunc("/", pkg.Home)
	http.HandleFunc("/about", pkg.About)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}

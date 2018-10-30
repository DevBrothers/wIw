package main

import (
	"fmt"
	"log"
	"net/http"
	"watIwant/apiserver"
)

type WatIWant struct{}

func (watIWant WatIWant) Run() {
	fmt.Println("watIwant service listening at port: 8080")
	err := http.ListenAndServe(":8080", apiserver.Handlers())

	if(err != nil){
		log.Fatal(":8080", err)
	}
}

func main()  {
	println("Hello World")
	watIwant := new(WatIWant)
	watIwant.Run()
}

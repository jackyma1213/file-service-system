package main

import (
	"file-service/controller"
	"fmt"
	"net/http"
)

func main() {

	port := "3000"
	route := controller.Init()
	fmt.Println("Server Started, listing on port " + port)
	http.ListenAndServe(":"+port, route)
}

package main

import (
	router "github.com/NirmalVatsyayan/FuzzyWuzzyGo/Router"
)


func StartServer(){
	port := ":8000"
	routes := router.Routes(port)
	routes.Run(port)
}

/*
check lsb_release -a on centos
env GOOS=linux GOARCH=amd64 go build -o "FuzzyMatch"
 */

func main(){
	StartServer()
}

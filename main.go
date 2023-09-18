package main

import "github.com/zsithh/go-rest-starter/cmd/app"

func main() {
	var a app.App
	err := a.Init()

	if err != nil {
		panic(err)
	}
	a.Run()
}

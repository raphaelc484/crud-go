package main

import (
	"crud-go/config"
	"crud-go/router"

	_ "crud-go/docs"
)

// @Title Crud Service API
func main() {
	err := config.Init()
	if err != nil {
		return
	}

	router.Initialize()
}

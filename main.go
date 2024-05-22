package main

import (
	"fmt"
	"ginStudy/config"
	"ginStudy/router"
)

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	r := router.InitRouter()

	port := config.Conf.Server.Port
	fmt.Println("Server is running on port:", port)
	r.Run(":8080")
}

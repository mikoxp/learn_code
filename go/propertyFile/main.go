package main

import (
	"fmt"

	"com.moles/loader"
)

func main() {
	config, err := loader.ReadConfig("test.properties")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Config:", config)
}

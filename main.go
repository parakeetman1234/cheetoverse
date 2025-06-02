package main

import (
	"fmt"
)

func main() {
	penis, err := LoadConfig("example_config.ini")
	if err != nil {
		panic(err)
	}
	fmt.Println(penis)
}

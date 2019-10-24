package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello world!")
	fmt.Println("api token: ", os.Getenv("SENSU_API_TOKEN"))
}

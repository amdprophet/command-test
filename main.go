package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello world!")

	fmt.Println("api token:", os.Getenv("SENSU_API_TOKEN"))
	fmt.Println("api refresh token:", os.Getenv("SENSU_API_REFRESH_TOKEN"))
	fmt.Println("api endpoints:", os.Getenv("SENSU_API_ENDPOINTS"))

	fmt.Println("path:", os.Getenv("PATH"))
	fmt.Println("ld_library_path:", os.Getenv("LD_LIBRARY_PATH"))
	fmt.Println("cpath:", os.Getenv("CPATH"))
}

package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello world!")

	fmt.Println("SENSU_API_URL=", os.Getenv("SENSU_API_URL"))
	fmt.Println("SENSU_NAMESPACE=", os.Getenv("SENSU_NAMESPACE"))
	fmt.Println("SENSU_FORMAT=", os.Getenv("SENSU_FORMAT"))
	fmt.Println("SENSU_ACCESS_TOKEN=", os.Getenv("SENSU_ACCESS_TOKEN"))
	fmt.Println("SENSU_ACCESS_TOKEN_EXPIRES_AT=", os.Getenv("SENSU_ACCESS_TOKEN_EXPIRES_AT"))
	fmt.Println("SENSU_REFRESH_TOKEN=", os.Getenv("SENSU_REFRESH_TOKEN"))
	fmt.Println("SENSU_TRUSTED_CA_FILE=", os.Getenv("SENSU_TRUSTED_CA_FILE"))
	fmt.Println("SENSU_INSECURE_SKIP_TLS_VERIFY=", os.Getenv("SENSU_INSECURE_SKIP_TLS_VERIFY"))

	fmt.Println("PATH=", os.Getenv("PATH"))
	fmt.Println("LD_LIBRARY_PATH=", os.Getenv("LD_LIBRARY_PATH"))
	fmt.Println("CPATH=", os.Getenv("CPATH"))
}

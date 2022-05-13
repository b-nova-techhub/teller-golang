package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Hello, %s", os.Getenv("ADMIN_USER"))
}

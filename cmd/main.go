package main

import (
	"fmt"
	"os"
)

func main() {
	user := "user"

	fmt.Printf("Hello %s, this is butter :)))\n", user)
	ReplStart(os.Stdin, os.Stdout)
}

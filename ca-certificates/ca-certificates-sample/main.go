package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("USAGE: ca-certificates <url>")
		os.Exit(1)
	}
	_, err := http.Head(os.Args[1])
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(2)
	}
	fmt.Println("SUCCESS!")
}

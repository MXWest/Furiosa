package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println("capable is the furiosa command line")
	args := os.Args[1:]
	if len(args) <= 0 {
		fmt.Printf("Error no arguments given!")
		os.Exit(1)
	} else {
		fmt.Printf("args is %v, len is %d\n", args, len(args))
	}
	furiosaHost := args[0]
	fmt.Printf("Host is %v", furiosaHost)

	endPoint := "http://" + furiosaHost + "/health"
	fmt.Printf("endPoint is %v\n", endPoint)

	resp, err := http.Get(endPoint)
	if err != nil {
		fmt.Printf("Problem contacting %v (%v)\n", endPoint, err)
		os.Exit(2)
	}

	fmt.Printf("Response is %v\n", resp)
}

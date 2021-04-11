package main

import (
	"casino_royal/vault/client"
	"flag"
	"fmt"
	"os"
)

var (
	backendURIFlag = flag.String("backend", "http://localhost:8080", "Backend API URL")
	helpFlag       = flag.Bool("help", false, "Display a helpful message.")
)

func main() {
	flag.Parse()
	s, err := client.NewSwitch()

	if err != nil {
		fmt.Printf("Cmd switch error: %s \n", err)
		os.Exit(2)
	}

	if *helpFlag || len(os.Args) == 1 {
		s.Help()
		return
	}

	err = s.Switch()
	if err != nil {
		fmt.Printf("Cmd switch error: %s \n", err)
		os.Exit(2)
	}

}

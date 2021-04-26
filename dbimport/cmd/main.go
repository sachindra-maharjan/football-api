package main

import (
	"casino_royal/vault/dbimport/client"
	"flag"
	"fmt"
	"os"
)

var (
	helpFlag = flag.Bool("help", false, "Display a helpful message.")
)

func main() {
	flag.Parse()
	s, err := client.NewSwitch()

	if err != nil {
		fmt.Printf("Cmd switch error: %s \n", err)
		os.Exit(2)
	}

	if *helpFlag || len(os.Args) < 3 {
		s.Help()
		return
	}

	err = s.Switch()
	if err != nil {
		fmt.Printf("Cmd switch error: %s \n", err)
		os.Exit(2)
	}

}

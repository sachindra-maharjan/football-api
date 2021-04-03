package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No command provided.")
		os.Exit(2)
	}

	cmd := os.Args[1]

	switch cmd {
	case "greet":
		greetCmd := flag.NewFlagSet("greet", flag.ExitOnError)
		msgFlag := greetCmd.String("msg", "Go CLI", "the greet message")
		err := greetCmd.Parse(os.Args[2:])
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Printf("Hello %s \n", *msgFlag)

	case "help":
		fmt.Println("help message")

	default:
		fmt.Printf("unknown command: %s \n", cmd)

	}

}

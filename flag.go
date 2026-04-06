package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Define a flag with a name, default value, and description
	// The return value is a pointer to the variable
	taskPtr := flag.String("add", "", "Add a new task description")
	priorityPtr := flag.Int("priority", 1, "Set task priority (1-5)")
	// Required to parse the arguments from os.Args[1:]
	flag.Parse()
	if *taskPtr != "" {
		fmt.Printf("Adding task: %s with priority: %d\n", *taskPtr, *priorityPtr)
	} else {
		fmt.Println("No task provided. Use -add \"task description\"")
		os.Exit(1)
	}
}

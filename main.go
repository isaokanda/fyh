package main

import (
	"os"
	"fmt"
	"path/filepath"
)

var version string

func main() {

	// version printing
	if len(os.Args) > 1 {
		if os.Args[1] == "--version" {
			fmt.Println(filepath.Base(os.Args[0])+" "+version)
		}
	}

	fmt.Println("Bon jour")
	fmt.Println("Ca va?")
	fmt.Println("Oui, ca va.")
	fmt.Println("Au revoir.")

	// loop test
	for i := 0; i < 5; i++ {
		fmt.Printf("%d\n", i)
	}

}

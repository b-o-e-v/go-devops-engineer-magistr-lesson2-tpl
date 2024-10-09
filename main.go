package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/b-o-e-v/go-devops-engineer-magistr-lesson2-tpl/validator"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: yamlvalid <filename>")
		os.Exit(1)
	}

	filename := os.Args[1]

	dirname := filepath.Dir(filename)
	relFilename, err := filepath.Rel(dirname, filename)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if errs := validator.Run(relFilename); len(errs) != 0 {
		for _, err := range errs {
			fmt.Println(err)
		}
		os.Exit(1)
	}

	fmt.Println("Validation successful")
	os.Exit(0)
}

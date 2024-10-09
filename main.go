package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/b-o-e-v/go-devops-engineer-magistr-lesson2-tpl/validator"
)

func main() {
	if len(os.Args) < 2 {
		panic("usage: yamlvalid <filename>")
	}

	filename := os.Args[1]

	_, err := os.Stat(filename)

	if errors.Is(err, os.ErrNotExist) {
		panic(fmt.Sprintf("%s does not exist", filename))
	}

	dirname := filepath.Dir(filename)
	relFilename, _ := filepath.Rel(dirname, filename)

	if errs := validator.Run(relFilename); len(errs) != 0 {
		for _, err := range errs {
			fmt.Println(err)
		}
	}
}

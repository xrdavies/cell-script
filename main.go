package main

import (
	"log"
	"os"

	"github.com/cell-labs/cell-script/internal/compiler"
)

func init() {
	compiler.SetupOptions()
}

func main() {
	err := compiler.Run(compiler.ParseOptions())
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

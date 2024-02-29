package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/wc-tool/args"
	"github.com/wc-tool/calc"
)

func main() {
	flags := args.ReadFlag()
	fileName := flag.Arg(0)

	fileData, err := args.GetFileContents(fileName)

	if err != nil {
		fmt.Printf("Error while reading the file: %v\n", err)
		os.Exit(1)
	}

	counts := calc.GetAllCounts(fileData)
	calc.PrintCounts(flags, counts, fileName)
}

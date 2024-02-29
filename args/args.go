package args

import (
	"flag"
	"io"
	"os"
)

type Args struct {
	CountBytes bool
	CountLines bool
	CountWords bool
	CountChars bool
}

func ReadFlag() *Args {
	var args Args
	flag.BoolVar(&args.CountBytes, "c", false, "Count Bytes")
	flag.BoolVar(&args.CountLines, "l", false, "Count Lines")
	flag.BoolVar(&args.CountWords, "w", false, "Count Words")
	flag.BoolVar(&args.CountChars, "m", false, "Count Characters")
	flag.Parse()
	return &args
}

func GetFileContents(fileName string) ([]byte, error) {
	if fileName == "" {
		return io.ReadAll(os.Stdin)
	}

	return os.ReadFile(fileName)
}

package main

import (
	"flag"

	"github.com/Mayowa-Ojo/exenv/cmd"
)

var (
	DEFAULT_LOCATION  = ".env"
	DEFAULT_DELIMITER = "="
)

func main() {
	delimiter := flag.String("sep", DEFAULT_DELIMITER, "Your key-value pair separator")
	location := flag.String("path", DEFAULT_LOCATION, "Your env file location. Relative/absolute")

	flag.Parse()

	if err := cmd.Execute(*location, *delimiter); err != nil {
		panic(err)
	}
}

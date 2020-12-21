package main

import (
	"fmt"
	"log"
	"os"
)

const (
	DefaultBrandNameCfgPath = "config/brandcfg/brand.yaml"
)

func main() {
	if err := run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(args []string) error {
	fmt.Print("Hello world")
	return nil
}

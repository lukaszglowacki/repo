package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/lukaszglowacki/repo/pkg/repo"
)

func main() {
	flag.Parse()
	command := repo.NewDefaultRepoCommand()

	if err := command.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

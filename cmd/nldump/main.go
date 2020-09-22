package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/slankdev/nlgo/internal/app/nldump"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	command := nldump.NewCommand()
	err := command.Execute()
	if err != nil {
		os.Exit(1)
	}
}

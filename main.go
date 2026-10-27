package main

import (
	"fmt"
	"os"
	"time"

	"github.com/pterm/pterm"
)

func main() {
	c := ViewConfig()
	time.Sleep(time.Second * 1)
	fmt.Print("\033[3;J\033[H\033[2J")
	args := os.Args
	switch args[1] {
	case "-c":
		Create(args[2], c)
	case "-r":
		Remove(c)
	default:
		pterm.Warning.Print("Tienes que elegir una opcion disponible. [-c o -r]")
		return
	}
}

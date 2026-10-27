package main

import (
	"os"

	"github.com/pterm/pterm"
)

/*
pedir un icono
pedir o conseguir un mimetype
dar chance a un Exec personalizado
pedir comentario
pedir tipo
pedir una categoria
*/

func main() {
	args := os.Args
	switch args[1] {
	case "-c":
		Create(args[2])
	case "-r":
		Remove()
	default:
		pterm.Warning.Print("Tienes que elegir una opcion disponible. [-c o -r]")
		return
	}
}

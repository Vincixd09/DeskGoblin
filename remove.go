package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/pterm/pterm"
)

var (
	options []string
)

func Remove(config Config) {
	filepath.Walk(config.Paths.Desktopsdir, func(path string, info fs.FileInfo, err error) error {

		if !info.IsDir() {
			options = append(options, path)
		}

		if err != nil {
			return err
		}

		return nil
	})
	options = append(options, "Back")
	desktop(options)

}

func desktop(options []string) {
	// Crear el componente multiselect interactivo
	selectedOptions, err := pterm.DefaultInteractiveMultiselect.
		WithOptions(options).
		WithFilter(true).
		Show("Selecciona el/los .desktop que quieras borrar")

	if err != nil {
		pterm.Fatal.Println("Error al capturar la selección:", err)
		desktop(options)
	}

	// Si el usuario no seleccionó nada
	if len(selectedOptions) == 0 {
		pterm.Warning.Println("No seleccionaste ningun .desktop.")
	}

	if selectedOptions[0] != "Back" {
		result, _ := pterm.DefaultInteractiveConfirm.WithDefaultText("Seguro/a que quieres eliminar esto?").Show()
		if result {
			fmt.Println(pterm.Green("Yes"))
			for i := range selectedOptions {
				err := os.Remove(selectedOptions[i])
				if err != nil {
					pterm.Error.Print(err)
					return
				}
			}

			pterm.Println() // Espacio en blanco decorativo
			pterm.Success.Println("¡.desktop borrado correctamente!")
			return
		}
		fmt.Println(pterm.Red("No"))
	}
}

package main

import (
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/pterm/pterm"
)

type Config struct {
	Paths struct {
		Desktopsdir string `toml:"desktopsdir"`
	} `toml:"path"`
}

func LoadConfig() string {
	configDir, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}

	appDir := filepath.Join(configDir, "DeskGoblin")

	err = os.MkdirAll(appDir, 0755)
	if err != nil {
		pterm.Error.Printfln("Error al crear la carpeta de la configuracion: %v", err)
		return ""
	}

	configFile := filepath.Join(appDir, "config.toml")

	// Crear el archivo solo si todavía no existe.
	file, err := os.OpenFile(
		configFile,
		os.O_CREATE|os.O_EXCL|os.O_WRONLY,
		0644,
	)
	if err != nil {
		if os.IsExist(err) {
			pterm.Info.Printfln("La configuración ya existe: %v", err)
			return configFile
		}
		panic(err)
	}
	defer file.Close()

	_, err = file.WriteString(`[path]
desktopsdir = "/home/kalisofi/.local/share/applications/"
`)
	if err != nil {
		panic(err)
	}

	return configFile
}

func ViewConfig() Config {
	var config Config
	path := LoadConfig()

	_, err := toml.DecodeFile(path, &config)
	if err != nil {
		pterm.Error.Printf("Error al ver la configuracion: %v", err)
		return config
	}

	return config
}

//Example

// go get github.com/BurntSushi/toml
// package main

// import (
//     "fmt"
//     "github.com/BurntSushi/toml"
// )

// type Config struct {
//     App struct {
//         Name    string `toml:"name"`
//         Version string `toml:"version"`
//     } `toml:"app"`

//     Paths struct {
//         JavaDir     string `toml:"java_dir"`
//         DownloadDir string `toml:"download_dir"`
//     } `toml:"paths"`

//     Settings struct {
//         ConfirmBeforeDelete bool   `toml:"confirm_before_delete"`
//         Theme               string `toml:"theme"`
//     } `toml:"settings"`
// }

// func main() {
//     var config Config

//     _, err := toml.DecodeFile("config.toml", &config)
//     if err != nil {
//         panic(err)
//     }

//     fmt.Println("Aplicación:", config.App.Name)
//     fmt.Println("Directorio de Java:", config.Paths.JavaDir)
//     fmt.Println("Tema:", config.Settings.Theme)
// }

// /home/tu_usuario/.config.
// package main

// import (
// 	"fmt"
// 	"os"
// 	"path/filepath"
// )

// func main() {
// 	configDir, err := os.UserConfigDir()
// 	if err != nil {
// 		panic(err)
// 	}

// 	appDir := filepath.Join(configDir, "JavaOptionsCli")

// 	err = os.MkdirAll(appDir, 0755)
// 	if err != nil {
// 		panic(err)
// 	}

// 	configFile := filepath.Join(appDir, "config.toml")

// 	// Crear el archivo solo si todavía no existe.
// 	file, err := os.OpenFile(
// 		configFile,
// 		os.O_CREATE|os.O_EXCL|os.O_WRONLY,
// 		0644,
// 	)
// 	if err != nil {
// 		if os.IsExist(err) {
// 			fmt.Println("La configuración ya existe:", configFile)
// 			return
// 		}
// 		panic(err)
// 	}
// 	defer file.Close()

// 	_, err = file.WriteString(`[settings]
// theme = "dark"
// `)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("Configuración creada:", configFile)
// }
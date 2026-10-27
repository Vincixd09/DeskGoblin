package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/pterm/pterm"
	"github.com/sqweek/dialog"
)

func Create(file string) {
	pterm.DefaultBox.WithRightPadding(10).WithLeftPadding(10).WithTopPadding(2).WithBottomPadding(2).Println("Holiwis :3")

	name, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("Ingrese un nombre").Show()
	comment, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("Ingrese un comentario").Show()
	t := terminal()
	e := exec(file)
	i := icon()
	m := mimetype(file)
	c := categories()

	desktop := fmt.Sprintf(`[Desktop Entry]
Name=%v
Comment=%v
Terminal=%v
Exec=%v
Type=Application
MimeType=%v
%v
%v`, name, comment, t, e, m, c, i)

	fmt.Println(desktop)
	err := os.Chdir("/home/kalisofi/.local/share/applications/")
	if err != nil {
		fmt.Println(err)
	}
	os.WriteFile(name+".desktop", []byte(desktop), 0755)
}

func terminal() string {
	result, _ := pterm.DefaultInteractiveConfirm.WithDefaultText("Queres que se abra la terminal al ejecutar?").Show()
	if result {
		fmt.Println(pterm.Green("Yes"))
		return "true"
	}
	fmt.Println(pterm.Red("No"))
	return "false"
}

func exec(appName string) string {
	var (
		r         string
		execPerso string
	)
	result, _ := pterm.DefaultInteractiveConfirm.WithDefaultText("Utilizar un exec personalizado?").Show()
	if result {
		result2, _ := pterm.DefaultInteractiveConfirm.WithDefaultText("Quieres utilizar un selector de archivos para buscar algun programa?").Show()
		if result2 {
			fmt.Println(pterm.Green("Yes"))
			r = dialogFn()
			result3, _ := pterm.DefaultInteractiveConfirm.WithDefaultText("Quieres agregar algo a la orden(Ej. [/usr/local/bin/alacritty] -> asi esta por defecto a selecionar | [/usr/local/bin/alacritty -e] -> aca agregas algo a la orden?").Show()
			if result3 {
				fmt.Println(pterm.Green("Yes"))
				execPerso, _ = pterm.DefaultInteractiveTextInput.WithDefaultText(fmt.Sprintf("Agregar algo a la orden: %v", r)).Show()
				execPerso2, _ := pterm.DefaultInteractiveTextInput.WithDefaultText(fmt.Sprintf("Agregar algo a la orden: %v", appName)).Show()
				return r + fmt.Sprintf(` %v "%v" %v`, execPerso, appName, execPerso2)
			}
			fmt.Println(pterm.Red("No"))
			return r + fmt.Sprintf(` "%v"`, appName)
		}
		fmt.Println("Tenga en cuenta que tambien tiene que agregar la ruta de su app en caso de que agregue otra, por ejemplo si agrega [/usr/local/bin/alacritty -e ] despues de eso tiene que agregar el path de su app y el argumento que vaya a usar, si no va a fallar")
		execPerso, _ = pterm.DefaultInteractiveTextInput.WithDefaultText("Ingrese la orden").Show()
		return execPerso
	}
	fmt.Println(pterm.Red("No"))
	return fmt.Sprintf(`"%v"`, appName)
}

func icon() string {
	result, _ := pterm.DefaultInteractiveConfirm.WithDefaultText("Utilizar icono?").Show()
	if result {
		fmt.Println(pterm.Green("Yes"))
		iconPath, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("Ingrese una ruta del icono").Show()
		return "Icon=" + iconPath
	}
	return ""
}

func mimetype(appName string) string {
	file, err := os.Open(appName)
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return ""
	}
	defer file.Close()

	// 2. Crear un buffer para los primeros 512 bytes
	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		fmt.Println("Error al leer el archivo:", err)
		return ""
	}

	// 3. Detectar el MIME type pasando el buffer
	return http.DetectContentType(buffer)
}

func categories() string {
	fmt.Println("Selector de Categorías para .desktop (Debian 13)")

	// Listado de Categorías Principales oficiales según la especificación de FreeDesktop
	options := []string{
		"AudioVideo      - Aplicación multimedia (Audio y/o Video)",
		"Audio           - Aplicación de sonido (Requiere también AudioVideo)",
		"Video           - Aplicación de video (Requiere también AudioVideo)",
		"Development     - Herramientas de desarrollo, IDEs y programación",
		"Education       - Software educativo y de aprendizaje",
		"Game            - Videojuegos y entretenimiento",
		"Graphics        - Aplicaciones de diseño gráfico y edición de imagen",
		"Network         - Software de red, navegadores Web, clientes de correo",
		"Office          - Suites de oficina, editores de texto y hojas de cálculo",
		"Science         - Software científico y matemático",
		"Settings        - Aplicaciones de configuración del sistema",
		"System          - Herramientas del sistema, monitores o administración",
		"Utility         - Utilidades y accesorios esenciales (Calculadoras, etc.)",
	}

	// Crear el componente multiselect interactivo
	selectedOptions, err := pterm.DefaultInteractiveMultiselect.
		WithOptions(options).
		WithFilter(true). // Permite al usuario escribir para buscar/filtrar en tiempo real
		Show("Selecciona las categorías para tu archivo .desktop (Espacio para marcar, Enter para confirmar):")

	if err != nil {
		pterm.Fatal.Println("Error al capturar la selección:", err)
		return ""
	}

	// Si el usuario no seleccionó nada
	if len(selectedOptions) == 0 {
		pterm.Warning.Println("No seleccionaste ninguna categoría.")
		return categories()
	}

	// Procesar las respuestas para extraer únicamente la clave limpia (ej: "AudioVideo")
	var finalCategories []string
	for _, option := range selectedOptions {
		// Divide por el espacio/guión y se queda con el primer elemento
		cleanCategory := strings.Fields(option)[0]
		finalCategories = append(finalCategories, cleanCategory)
	}

	// Unir las categorías mediante punto y coma (;) como exige el formato de los archivos .desktop
	categoriesString := strings.Join(finalCategories, ";") + ";"

	pterm.Println() // Espacio en blanco decorativo
	pterm.Success.Println("¡Categorías procesadas correctamente!")

	// Mostrar el resultado formateado listo para copiar en el .desktop
	pterm.DefaultSection.Println("Resultado para tu archivo .desktop")
	return fmt.Sprintf("Categories=%s", categoriesString)
}

func dialogFn() string {
	f, err := dialog.File().Title("Elija una carpeta").SetStartDir("/usr/local/bin").Load()
	if err != nil {
		dialog.Message("Error: %v\n", err).Title("Error")
		return ""
	}
	return f
}

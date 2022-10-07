package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/colors", getColors)
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		return
	}

}

func getColors(writer http.ResponseWriter, request *http.Request) {
	tmpl, _ := template.ParseFiles("templates/layout.html")
	rgb := request.URL.Query().Get("rgb")
	colors := initColorsNamesSet()
	fmt.Println(colors)

	colorName := getColorNameByRgb(rgb, colors)

	page := Page{PageTitle: colorName, ColorRGB: template.CSS(rgb)}
	err := tmpl.Execute(writer, page)
	if err != nil {
		return
	}
}

type Page struct {
	PageTitle string
	ColorRGB  template.CSS
}

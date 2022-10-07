package main

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"net/http"
)

func main() {
	//db := getDbConnection()
	//initDb(db)

	http.HandleFunc("/colors", getColors)
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		return
	}

	//err = db.Close()
	//if err != nil {
	//	return
	//}
}

func getColors(writer http.ResponseWriter, request *http.Request) {
	tmpl, _ := template.ParseFiles("templates/layout.html")
	rgb := request.URL.Query().Get("rgb")
	colors := initColorsNamesSet()
	fmt.Println(colors)
	//db, err := sql.Open("sqlite3", "colors.db")
	//db := getDbConnection()
	//initDb(db)
	colorName := getColorNameByRgb(rgb, colors)
	//page := Page{PageTitle: "test page title", ColorRGB: template.CSS(rgb)}
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

type Color struct {
	rgb  string
	name string
}

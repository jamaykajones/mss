package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/jamaykajones/mss/webapp/controller"
)

func main() {
	templates := populateTemplates()
	controller.Startup(templates)
	http.ListenAndServe(":8000", nil)
}

func populateTemplates() map[string]*template.Template { // returns a single template back out
	result := make(map[string]*template.Template)
	const basePath = "templates"
	layout := template.Must(template.ParseFiles(basePath + "/_layout.html"))
	template.Must(layout.ParseFiles(basePath+"/_header.html", basePath+"/_footer.html"))
	dir, err := os.Open(basePath + "/content")
	if err != nil {
		panic("Failed to open template blocks dir: " + err.Error())
	}
	fis, err := dir.Readdir(-1)
	if err != nil {
		panic("Failed to read contents of content dir: " + err.Error())
	}
	for _, fileInfo := range fis {
		f, err := os.Open(basePath + "/content/" + fileInfo.Name())
		if err != nil {
			panic("Failed to open template '" + fileInfo.Name() + "'")
		}
		content, err := ioutil.ReadAll(f)
		if err != nil {
			panic("Failed to read content from file '" + fileInfo.Name() + "'")
		}
		f.Close()
		tmpl := template.Must(layout.Clone())
		_, err = tmpl.Parse(string(content))
		if err != nil {
			panic("Failed to parse contents of '" + fileInfo.Name() + "' as template")
		}
		result[fileInfo.Name()] = tmpl
	}
	return result
}

package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("template.gohtml"))
}

type animal struct {
	Name string
	Kind string
	Age  int
}

type templateData struct {
	Td     []string
	Lang   map[string]string
	Animal animal
}

func main() {
	todos := []string{"write", "meditate", "code", "read"}
	languages := map[string]string{
		"Japan":   "Japanese",
		"England": "English",
		"Mexico":  "Spanish",
		"China":   "Chinese",
	}
	oreo := animal{
		Name: "Oreo",
		Kind: "Dog",
		Age:  4,
	}

	tplData := templateData{
		Td:     todos,
		Lang:   languages,
		Animal: oreo,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "template.gohtml", tplData)

	if err != nil {
		log.Fatal("Error executing template:", err)
	}
}

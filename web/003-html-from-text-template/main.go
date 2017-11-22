package main

import (
	"html/template"
	"log"
	"os"
)

// run this file to generate an index.html from the go template
func main() {
	tpl, err := template.ParseFiles("template.gohtml")
	if err != nil {
		log.Fatal("Error parsing the template", err)
	}

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("Error creating a file:", err)
	}

	defer nf.Close()

	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatal("Error executing the template:", err)
	}
}

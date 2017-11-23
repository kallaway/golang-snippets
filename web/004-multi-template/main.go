package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

// When you run this file, it takes multiple templates and loads them at once.
// Then we manually choose which one to print out
// Parsing the templates in init() enables us
// to only parse them once and improves efficiency
var tpl *template.Template

func init() {
	// Must() will panic automatically if it gets non-nil error
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatal("Error printing first template to stdout:", err)
	}

	fmt.Println()

	err = tpl.ExecuteTemplate(os.Stdout, "second.gohtml", nil)
	if err != nil {
		log.Fatal("Error printing second template to stdout:", err)
	}

	fmt.Println()

	err = tpl.ExecuteTemplate(os.Stdout, "third.gohtml", nil)
	if err != nil {
		log.Fatal("Error printing third template to stdout:", err)
	}
}

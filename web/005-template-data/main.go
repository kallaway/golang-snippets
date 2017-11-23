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

func main() {
	nums := []int{32, 14, 34}
	err := tpl.ExecuteTemplate(os.Stdout, "template.gohtml", nums)
	if err != nil {
		log.Fatal("Error executing template:", err)
	}
}

package main

import (
	"html/template"
	"log"
	"os"
	"strings"
)

var tpl *template.Template

var funcMap = template.FuncMap{
	"uc":     strings.ToUpper,
	"lc":     strings.ToLower,
	"limit3": displayFirstThree,
}

type college struct {
	Subjects []string
	Name     string
}

func init() {
	tpl = template.Must(template.New("").Funcs(funcMap).ParseFiles("template.gohtml"))
}

func displayFirstThree(items []string) string {
	itemsToShow := items[:3]
	listString := strings.Join(itemsToShow, ", ")

	return listString
}

func main() {

	studiedSubjects := []string{
		"Calculus",
		"English",
		"Physics",
		"Biology",
		"Economics",
		"Business Strategy",
	}

	collegeName := "Imaginary College of Mathematics & Technology"

	exampleCollege := college{
		studiedSubjects,
		collegeName,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "template.gohtml", exampleCollege)
	if err != nil {
		log.Fatalln("Error executing the template:", err)
	}
}

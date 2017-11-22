package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// expects 2 command line arguments
// First name and Last name

// It creates an index.html file with
// content filled in based on the arguments given
func main() {
	if len(os.Args) != 3 {
		log.Fatal(`Error: You haven't provided enough arguments. 
			This program expects 2 command line arguments: first name and last name.`)
	}

	fname := os.Args[1]
	lname := os.Args[2]

	template := fmt.Sprint(`<!DOCTYPE html>
	<html>
		<head>
			<meta charset="utf-8">
			<title></title>
		</head>
		<body>
			<h1>The name is ` + fname + `</h1>
			<h2>What's the last name? It's ` + lname + `</h2>
		</body>
	</html>`)

	nf, err := os.Create("index.html")

	if err != nil {
		log.Fatal("Error printing the file:", err)
	}

	defer nf.Close()

	io.Copy(nf, strings.NewReader(template))
}

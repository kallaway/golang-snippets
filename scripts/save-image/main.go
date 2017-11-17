package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://picsum.photos/200/300/")

	if err != nil {
		log.Fatal("Error:", err)
	}

	defer resp.Body.Close()

	file, err := os.Create("img/image1.jpg")
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(file, resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	file.Close()
}

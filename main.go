package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Config struct {
	Port      int      `json:"port"`
	FilePaths []string `json:"file_paths"`
}

var config = new(Config)

func main() {
	// Read config
	configFile, err := os.Open("ng-html5mode-serve.json")
	if err != nil {
		log.Fatalln(err)
	}

	configDecoder := json.NewDecoder(configFile)
	err = configDecoder.Decode(config)
	if err != nil {
		log.Fatalln(err)
	}

	// Handle file serve directories. Otherwise, serve index.html so Angular routes work.
	for _, filePath := range config.FilePaths {
		http.Handle(fmt.Sprintf("/%s/", filePath), http.FileServer(http.Dir("")))
	}
	http.HandleFunc("/", alwaysServeIndex)

	log.Println("Will serve on port", config.Port)

	err = http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil)
	log.Fatalln(err)
}

func alwaysServeIndex(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("index.html")
	defer file.Close()
	if err != nil {
		log.Fatalln(err)
	}

	_, err = io.Copy(w, file)
	if err != nil {
		log.Fatalln(err)
	}
}

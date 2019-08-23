package main

import (
	"log"
	"net/http"
	"./com/philoenglish/file"
	"./com/philoenglish/props"
)

const configFile = "config.properties"

func main() {

	if props, err := propsReader.ReadPropertiesFile(configFile); err != nil {
		log.Fatal("failed to load mandatory properties from ", configFile)
		panic(err)
	} else {
		port := props["httpPort"]
		log.Println("Starting server and listening on port ", port)
		http.HandleFunc("/", receiver.HomeHandler)
		http.HandleFunc("/upload", receiver.CreateHandler(props))
		log.Fatal(http.ListenAndServe(":"+port, nil))
	}
}
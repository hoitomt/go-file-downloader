package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	environment := "development"

	configfile := flag.String("c", "config.ini", "Location of the hudupdater config file [config.ini]")

	flag.Parse()

	log.Printf("using config file %s\n", *configfile)
	config := InitializeConfig(*configfile)

	// Print the parsed config file
	// requires github.com/davecgh/go-spew/spew
	// spew.Dump(config)

	extractPath := config.extractPath(environment)
	log.Printf("Extract Path for %s: %s", environment, extractPath)

	err := downloadFile(extractPath, config)
	if err != nil {
		panic(err)
	}
}

func downloadFile(filepath string, config *Config) (err error) {

	// Go has fancier ways to do this, but this is QAD for now
	url := config.BaseUri + "url/distribution?apikey=" + config.ApiKey + "&reports=true"
	log.Printf("curl %s\n", url)

	// Create the directory(filepath)
	err = os.MkdirAll(filepath, 0700)
	if err != nil {
		log.Println("Could not create the directory")
		return err
	}

	// Create the file
	out, err := os.Create(filepath + "/test.json")
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Could not retrieve the file")
		return err
	}
	defer resp.Body.Close()

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Println("Could not write the file")
		return err
	}

	return nil
}

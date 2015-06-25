package main

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

var SITES = []string{
	"https://www.google.com",
	"https://www.develer.com",
	"https://www.yahoo.com",
	"https://www.facebook.com",
}

func download(uri string) {
	resp, err := http.Get(uri)
	if err != nil {
		log.Println("error fetching file", uri, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		log.Println("http error:", uri, resp.Status)
		return
	}

	parsed, err := url.Parse(uri)
	if err != nil {
		log.Println("invalid URL", err)
		return
	}

	f, err := os.Create(parsed.Host)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	io.Copy(f, resp.Body)
	log.Println("downloaded:", uri)
}

func GetRobots() {
	for {
		for _, s := range SITES {
			go download(s + "/robots.txt")
		}
		time.Sleep(10 * time.Second)
	}
}

func main() {
	go GetRobots()
	select {}
}
package fetcher

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func download(uri string) error {
	resp, err := http.Get(uri)
	if err != nil {
		return fmt.Errorf("error fetching file: %v %v", uri, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("http error: %v %v", uri, resp.Status)
	}

	parsed, err := url.Parse(uri)
	if err != nil {
		return fmt.Errorf("invalid URL %v", err)
	}

	f, err := os.Create(parsed.Host)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return fmt.Errorf("error while fetching: %v", err)
	}
	log.Println("downloaded:", uri)
	return nil
}

func GetRobots() {
	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, syscall.SIGHUP)

	for {
		for _, s := range SITES {
			go func(s string) {
				err := download(s + "/robots.txt")
				if err != nil {
					log.Println("error downloading:", err)
				}
			}(s)
		}

		select {
		case <-sigch:
		case <-time.After(10 * time.Second):
		}
	}
}

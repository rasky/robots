package main

import "github.com/rasky/robots/fetcher"

func main() {
	go fetcher.GetRobots()
	select {}
}

package main

import "github.com/rasky/robots/fetcher"
import "github.com/Sirupsen/logrus"

func main() {
	logrus.Println("Ciao")
	go fetcher.GetRobots()
	select {}
}

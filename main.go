package main

import "log"

func main() {
	credential := parseCredentials()
	log.Println(credential)

	setupServer()
}

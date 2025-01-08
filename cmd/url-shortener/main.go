package main

import "url-shortener/cmd/servid"

func main() {
	server := servid.NewApp()
	server.Start()
}
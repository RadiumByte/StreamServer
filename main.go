package main

import (
	"github.com/RadiumByte/StreamServer/api"
	"github.com/RadiumByte/StreamServer/app"
	"github.com/RadiumByte/StreamServer/yal"

	"fmt"
)

func run(errc chan<- error) {
	youtubeClient, err := yal.NewYoutubeClient()
	if err != nil {
		errc <- err
		return
	}

	application := app.NewApplication(youtubeClient, errc)
	server := api.NewWebServer(application)

	server.Start(errc)
}

func main() {
	fmt.Println("Streaming Management Server is preparing to start...")

	errc := make(chan error)
	go run(errc)
	if err := <-errc; err != nil {
		fmt.Print("Error occured: ")
		fmt.Println(err)
	}
}

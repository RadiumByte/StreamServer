package main

import (
	"github.com/RadiumByte/StreamServer/api"
	"github.com/RadiumByte/StreamServer/app"
	"github.com/RadiumByte/StreamServer/ral"
	"github.com/RadiumByte/StreamServer/yal"

	"fmt"
)

func run(errc chan<- error) {
	CarIP := "192.168.183.50"
	Port := ":8080"

	robot, err := ral.NewRoboCar(CarIP, Port)
	if err != nil {
		errc <- err
		return
	}

	youtubeClient, err := yal.NewYoutubeClient()
	if err != nil {
		errc <- err
		return
	}

	application := app.NewApplication(youtubeClient, robot, errc)
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

package ral

import (
	"fmt"
	"strconv"

	"github.com/valyala/fasthttp"
)

// RoboCar represents Raspberry Pi based car
type RoboCar struct {
	Client   *fasthttp.Client
	Request  *fasthttp.Request
	Response *fasthttp.Response
	CarIP    string
	CarPort  string
}

// Turn sends only steering command to Car
func (robot *RoboCar) Turn(steerValue int) {
	if steerValue > 100 {
		steerValue = 100
	} else if steerValue < 0 {
		steerValue = 0
	}

	steerValueStr := strconv.Itoa(steerValue)
	command := "S" + steerValueStr + "A"
	//fmt.Println("Sending command: " + command)

	url := "http://" + robot.CarIP + robot.CarPort + "/" + command
	robot.Request.SetRequestURI(url)
	robot.Client.Do(robot.Request, robot.Response)

	//fmt.Println("Command sent to robot: " + command)
}

// DirectCommand sends any command, according to Car's specs
func (robot *RoboCar) DirectCommand(command string) {
	if command != "HALT" {
		command += "A"
	}

	fmt.Println("Sending command: " + command)

	url := "http://" + robot.CarIP + robot.CarPort + "/" + command
	robot.Request.SetRequestURI(url)
	robot.Client.Do(robot.Request, robot.Response)

	fmt.Println("Command sent to robot: " + command)
}

// NewRoboCar constructs object of RoboCar
func NewRoboCar(ip string, port string) (*RoboCar, error) {
	res := &RoboCar{}
	res.Client = &fasthttp.Client{}
	res.Request = fasthttp.AcquireRequest()
	res.Request.Header.SetMethod("PUT")

	res.Response = fasthttp.AcquireResponse()

	res.CarPort = port
	res.CarIP = ip
	return res, nil
}

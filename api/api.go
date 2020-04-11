package api

import (
	"encoding/json"
	"fmt"

	"github.com/RadiumByte/StreamServer/app"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

// GetCamerasJSON describes JSON data for returning camera data
type GetCamerasJSON struct {
	CameraTypes []int    `json:"types"`
	CameraNames []string `json:"names"`
}

// WebServer is responsible for communication with clients
type WebServer struct {
	application app.StreamManager
}

// GetCameras handles GET request for listing cameras
func (server *WebServer) GetCameras(ctx *fasthttp.RequestCtx) {
	fmt.Println("API: GET request /get-cameras accepted...")

	cameraList := server.application.GetCameras()

	if len(cameraList) != 0 {
		var types []int
		var names []string

		for i := range cameraList {
			types = append(types, cameraList[i].Type)
			names = append(names, cameraList[i].Name)
		}

		toEncode := &GetCamerasJSON{
			CameraTypes: types,
			CameraNames: names}

		payload, _ := json.Marshal(toEncode)
		fmt.Println("Server response for /get-cameras request: ")
		fmt.Println(string(payload))

		ctx.SetContentType("application/json")
		ctx.SetBodyString(string(payload))

		ctx.SetStatusCode(fasthttp.StatusOK)
	}

	ctx.SetStatusCode(fasthttp.StatusNoContent)
}

// SelectCamera handles POST request for selecting camera for streaming
func (server *WebServer) SelectCamera(ctx *fasthttp.RequestCtx) {
	fmt.Println("API: POST request /select-camera accepted...")

	// TO DO: get JSON from ctx body

	// TO DO: deserialize JSON to model

	// TO DO: find camera in App, if not found -> error code

	// TO DO: return status OK
}

// AddCamera handles POST request for adding new camera to the server's list
func (server *WebServer) AddCamera(ctx *fasthttp.RequestCtx) {
	fmt.Println("API: POST request /add-camera accepted...")

	// TO DO: get JSON from ctx body

	// TO DO: deserialize JSON to model

	// TO DO: create camera in App

	// TO DO: return status OK
}

// Start adds routes and begins serving
func (server *WebServer) Start(errc chan<- error) {
	router := fasthttprouter.New()

	// Routes for camera management
	router.GET("/get-cameras", server.GetCameras)
	router.POST("/select-camera", server.SelectCamera)
	router.POST("/add-camera", server.AddCamera)

	port := ":8081"

	fmt.Printf("Server is starting on port %s\n", port)
	errc <- fasthttp.ListenAndServe(port, router.Handler)
}

// NewWebServer constructs WebServer
func NewWebServer(application app.StreamManager) *WebServer {
	res := &WebServer{}
	res.application = application

	return res
}

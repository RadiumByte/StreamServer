package api

import (
	"encoding/json"
	"fmt"

	"github.com/RadiumByte/StreamServer/app"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

// GetCamerasJSON describes JSON data for returning multiple camera data
type GetCamerasJSON struct {
	CameraTypes []int    `json:"types"`
	CameraNames []string `json:"names"`
}

// GetActiveJSON describes JSON data for returning one active camera data
type GetActiveJSON struct {
	CameraType int    `json:"type"`
	CameraName string `json:"name"`
}

// WebServer is responsible for communication with clients
type WebServer struct {
	application app.StreamManager
}

// GetCameras handles GET request for listing cameras
func (server *WebServer) GetCameras(ctx *fasthttp.RequestCtx) {
	fmt.Println("-----------------------------------------")
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
		ctx.SetBody(payload)

		fmt.Println("Server status for /get-cameras request: OK")
		ctx.SetStatusCode(fasthttp.StatusOK)
		return
	}

	fmt.Println("Server status for /get-cameras request: NoContent")
	ctx.SetStatusCode(fasthttp.StatusNoContent)
}

// GetActive handles GET request for viewing current (active) camera
func (server *WebServer) GetActive(ctx *fasthttp.RequestCtx) {
	fmt.Println("-----------------------------------------")
	fmt.Println("API: GET request /get-active accepted...")

	camera := server.application.GetActive()

	if camera.Name != "" {
		toEncode := &GetActiveJSON{
			CameraType: camera.Type,
			CameraName: camera.Name}

		payload, _ := json.Marshal(toEncode)
		fmt.Println("Server response for /get-active request: ")
		fmt.Println(string(payload))

		ctx.SetContentType("application/json")
		ctx.SetBody(payload)

		fmt.Println("Server status for /get-active request: OK")
		ctx.SetStatusCode(fasthttp.StatusOK)
		return
	}

	fmt.Println("Server status for /get-active request: NoContent")
	ctx.SetStatusCode(fasthttp.StatusNoContent)
}

// SelectCamera handles POST request for selecting camera for streaming
func (server *WebServer) SelectCamera(ctx *fasthttp.RequestCtx) {
	fmt.Println("-----------------------------------------")
	fmt.Println("API: POST request /select-camera accepted...")

	payload := ctx.PostBody()

	var dataJSON map[string]interface{}
	if err := json.Unmarshal(payload, &dataJSON); err != nil {
		fmt.Println("Server status for /select-camera request: BadRequest")
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	name := dataJSON["name"].(string)
	fmt.Printf("Server got camera name: %s\n", name)

	err := server.application.SelectCamera(name)
	if err != nil {
		fmt.Println("Server status for /select-camera request: NoContent")
		ctx.SetStatusCode(fasthttp.StatusNoContent)
		return
	}
	fmt.Println("Server selected camera successfully")

	ctx.SetStatusCode(fasthttp.StatusOK)
}

// SetStreamURL handles POST request for saving broadcast URL
func (server *WebServer) SetStreamURL(ctx *fasthttp.RequestCtx) {
	fmt.Println("-----------------------------------------")
	fmt.Println("API: POST request /stream-url accepted...")

	URL := string(ctx.PostBody())

	fmt.Printf("Server got broadcast URL: %s\n", URL)

	server.application.SetStreamURL(URL)

	ctx.SetStatusCode(fasthttp.StatusOK)
}

// GetStreamURL handles GET request for taking broadcast URL
func (server *WebServer) GetStreamURL(ctx *fasthttp.RequestCtx) {
	fmt.Println("-----------------------------------------")
	fmt.Println("API: GET request /stream-url accepted...")

	URL := server.application.GetStreamURL()

	ctx.SetBodyString(URL)
	ctx.SetStatusCode(fasthttp.StatusOK)
}

// AddCamera handles POST request for adding new camera to the server's list
func (server *WebServer) AddCamera(ctx *fasthttp.RequestCtx) {
	fmt.Println("-----------------------------------------")
	fmt.Println("API: POST request /add-camera accepted...")

	payload := ctx.PostBody()
	var dataJSON map[string]interface{}
	if err := json.Unmarshal(payload, &dataJSON); err != nil {
		fmt.Println("Server status for /add-camera request: BadRequest")
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	nameCam := dataJSON["name"].(string)
	typeCam := dataJSON["type"].(float64)
	urlCam := dataJSON["url"].(string)
	fmt.Printf("Server got camera data: \n Name: %s\n Type: %d\n URL (or device): %s\n", nameCam, int(typeCam), urlCam)

	newCamera := app.CameraData{
		Name: nameCam,
		Type: int(typeCam),
		URL:  urlCam}

	server.application.AddCamera(newCamera)

	ctx.SetStatusCode(fasthttp.StatusOK)
}

// Start adds routes and begins serving
func (server *WebServer) Start(errc chan<- error) {
	router := fasthttprouter.New()

	// Routes for camera management
	router.GET("/get-cameras", server.GetCameras)
	router.POST("/select-camera", server.SelectCamera)
	router.POST("/add-camera", server.AddCamera)
	router.GET("/get-active", server.GetActive)
	router.POST("/stream-url", server.SetStreamURL)
	router.GET("/stream-url", server.GetStreamURL)

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

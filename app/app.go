package app

import (
	"errors"
	"sync"
)

// StreamManager is an interface for application's core
type StreamManager interface {
	AddCamera(camera CameraData)
	SelectCamera(name string) error
	GetCameras() []CameraData
	GetActive() CameraData
}

// YoutubeAccessLayer is an interface for calling Youtube from Application
type YoutubeAccessLayer interface {
	RunRTMP(rtspInput string, rtmpOutput string)
}

// Application is responsible for all logics and communicates with other layers
type Application struct {
	Youtube YoutubeAccessLayer
	errc    chan<- error

	// Storage for cameras
	cameras      []CameraData
	activeCamera CameraData

	camerasMutex sync.Mutex
}

// AddCamera creates new camera in list
func (a *Application) AddCamera(camera CameraData) {
	a.camerasMutex.Lock()
	defer a.camerasMutex.Unlock()

	a.cameras = append(a.cameras, camera)
}

// SelectCamera switches stream to specified camera
// If specified camera is unknown - returns error
func (a *Application) SelectCamera(name string) error {
	a.camerasMutex.Lock()
	defer a.camerasMutex.Unlock()

	for i := range a.cameras {
		if a.cameras[i].Name == name {
			a.activeCamera.Name = a.cameras[i].Name
			a.activeCamera.Type = a.cameras[i].Type
			a.activeCamera.URL = a.cameras[i].URL

			a.Youtube.RunRTMP(a.activeCamera.URL, "rtmp://a.rtmp.youtube.com/live2/uq80-f37c-z3c4-7rth")

			return nil
		}
	}
	return errors.New("Camera name is invalid")
}

// GetCameras returns list of added cameras
func (a *Application) GetCameras() []CameraData {
	return a.cameras
}

// GetActive returns one active camera
func (a *Application) GetActive() CameraData {
	return a.activeCamera
}

// NewApplication constructs Application
func NewApplication(youtube YoutubeAccessLayer, errchannel chan<- error) *Application {
	res := &Application{}

	res.Youtube = youtube
	res.errc = errchannel
	res.cameras = []CameraData{}
	res.activeCamera = CameraData{
		Name: "",
		URL:  "",
		Type: 0}

	return res
}

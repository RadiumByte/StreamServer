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
	SetStreamURL(URL string)
	GetStreamURL() string
}

// YoutubeAccessLayer is an interface for calling Youtube from Application
type YoutubeAccessLayer interface {
	RunIPRTMP(rtspInput string, cameraType int, rtmpOutput string)
	RunDevRTMP(deviceInput string, rtmpOutput string)
	RunVLC(deviceInput string, rtspOutput string)
}

// RobotAccessLayer is an interface for RAL usage from Application
type RobotAccessLayer interface {
	Turn(int)
	DirectCommand(string)
}

// Application is responsible for all logics and communicates with other layers
type Application struct {
	Robot   RobotAccessLayer
	Youtube YoutubeAccessLayer
	errc    chan<- error

	// Storage for cameras
	cameras      []CameraData
	activeCamera CameraData

	streamURL string

	camerasMutex sync.Mutex
}

// SetStreamURL sets new stream URL
func (a *Application) SetStreamURL(URL string) {
	a.streamURL = URL
}

// GetStreamURL gets new stream URL
func (a *Application) GetStreamURL() string {
	return a.streamURL
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

			cameraOutput := a.activeCamera.URL

			// If the camera is /dev/video
			if a.activeCamera.Type == 0 {
				/*OLD VLC TRANSCODING
				rtspOutput := "rtsp://localhost:8554/"
				a.Youtube.RunVLC(a.activeCamera.URL, rtspOutput)
				time.Sleep(2 * time.Second)
				a.Youtube.RunIPRTMP(rtspOutput, a.activeCamera.Type, "rtmp://a.rtmp.youtube.com/live2/uq80-f37c-z3c4-7rth")
				*/
				a.Youtube.RunDevRTMP(cameraOutput, "rtmp://a.rtmp.youtube.com/live2/uq80-f37c-z3c4-7rth")
			} else {
				a.Youtube.RunIPRTMP(cameraOutput, a.activeCamera.Type, "rtmp://a.rtmp.youtube.com/live2/uq80-f37c-z3c4-7rth")
			}

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
func NewApplication(youtube YoutubeAccessLayer, robot RobotAccessLayer, errchannel chan<- error) *Application {
	res := &Application{}

	res.Youtube = youtube
	res.errc = errchannel
	res.Robot = robot

	res.cameras = []CameraData{}
	res.activeCamera = CameraData{
		Name: "",
		URL:  "",
		Type: 0}

	return res
}

package app

// StreamManager is an interface for application's core
type StreamManager interface {
	AddCamera(camera CameraData)
	SelectCamera(name string) error
	GetCameras() []CameraData
}

// YoutubeAccessLayer is an interface for calling Youtube from Application
type YoutubeAccessLayer interface {
}

// Application is responsible for all logics and communicates with other layers
type Application struct {
	Youtube YoutubeAccessLayer
	errc    chan<- error

	// Storage for cameras
	cameras []CameraData
}

// AddCamera creates new camera in list
func (a *Application) AddCamera(camera CameraData) {
	return
}

// SelectCamera switches stream to specified camera
// If specified camera is unknown - returns error
func (a *Application) SelectCamera(name string) error {
	return nil
}

// GetCameras returns list of added cameras
func (a *Application) GetCameras() []CameraData {
	return []CameraData{}
}

// NewApplication constructs Application
func NewApplication(youtube YoutubeAccessLayer, errchannel chan<- error) *Application {
	res := &Application{}

	res.Youtube = youtube
	res.errc = errchannel
	res.cameras = []CameraData{}

	return res
}

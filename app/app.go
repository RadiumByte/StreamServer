package app

// StreamManager is an interface for application's core
type StreamManager interface {
}

// YoutubeAccessLayer is an interface for calling Youtube from Application
type YoutubeAccessLayer interface {
}

// Application is responsible for all logics and communicates with other layers
type Application struct {
	Youtube YoutubeAccessLayer
	errc    chan<- error
}

// NewApplication constructs Application
func NewApplication(youtube YoutubeAccessLayer, errchannel chan<- error) *Application {
	res := &Application{}

	res.Youtube = youtube
	res.errc = errchannel

	return res
}

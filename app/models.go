package app

// CameraData describes general info about cameras
type CameraData struct {
	Name string
	URL  string

	// 0 - is a /dev camera
	// 1 - is a TCP RTSP camera
	// 2 - is a UDP RTSP camera
	Type int
}

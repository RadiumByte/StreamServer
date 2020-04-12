package yal

import (
	"os/exec"
)

// YoutubeClient represents data for connection to Youtube
type YoutubeClient struct {
}

// RunRTMP coverts input RTSP stream to RTMP stream using FFMPEG
func (youtube *YoutubeClient) RunRTMP(rtspInput string, cameraType int, rtmpOutput string) {
	cmdKill := exec.Command("killall", "ffmpeg")
	cmdKill.Run()

	transport := ""

	if cameraType == 1 {
		transport = "tcp"
	} else if cameraType == 2 {
		transport = "udp"
	}

	cmdFFMPEG := exec.Command("ffmpeg", "-f", "lavfi", "-i", "anullsrc", "-rtsp_transport", transport, "-i", rtspInput, "-tune", "zerolatency", "-vcodec", "libx264", "-t", "12:00:00", "-pix_fmt", "+", "-c:v", "copy", "-c:a", "aac", "-strict", "experimental", "-f", "flv", rtmpOutput)
	cmdFFMPEG.Start()
}

// NewYoutubeClient constructs object of YoutubeClient
func NewYoutubeClient() (*YoutubeClient, error) {
	res := &YoutubeClient{}

	return res, nil
}

/*
ffmpeg -f lavfi -i anullsrc -rtsp_transport tcp -i rtsp://81.23.197.208/user=admin_password=8555_channel=4_stream=0.sdp -tune zerolatency -vcodec libx264 -t 12:00:00 -pix_fmt + -c:v copy -c:a aac -strict experimental -f flv rtmp://a.rtmp.youtube.com/live2/uq80-f37c-z3c4-7rth
*/

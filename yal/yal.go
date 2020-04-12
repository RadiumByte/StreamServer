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

	if cameraType != 0 {
		cmdKill = exec.Command("killall", "cvlc")
		cmdKill.Run()
	}

	transport := ""

	if cameraType == 1 {
		transport = "tcp"
	} else if cameraType == 2 || cameraType == 0 {
		transport = "udp"
	}

	cmdFFMPEG := exec.Command("ffmpeg", "-f", "lavfi", "-i", "anullsrc", "-rtsp_transport", transport, "-i", rtspInput, "-tune", "zerolatency", "-vcodec", "libx264", "-t", "12:00:00", "-pix_fmt", "+", "-c:v", "copy", "-c:a", "aac", "-strict", "experimental", "-f", "flv", rtmpOutput)
	cmdFFMPEG.Start()
}

//cvlc -vvv v4l2:///dev/video0 --sout '#transcode{vcodec=mp2v,vb=800,acodec=none}:rtp{sdp=rtsp://:8554/}'

// RunVLC coverts input /dev/video stream to RTSP stream using VLC
func (youtube *YoutubeClient) RunVLC(deviceInput string, rtspOutput string) {
	cmdKill := exec.Command("killall", "cvlc")
	cmdKill.Run()

	cmdKill = exec.Command("killall", "ffmpeg")
	cmdKill.Run()

	cmdVLC := exec.Command("cvlc", "-vvv", "v4l2://"+deviceInput, "--sout", "#transcode{vcodec=h264,vb=2000,acodec=none}:rtp{sdp="+rtspOutput+"}")
	cmdVLC.Start()
}

// NewYoutubeClient constructs object of YoutubeClient
func NewYoutubeClient() (*YoutubeClient, error) {
	res := &YoutubeClient{}

	return res, nil
}

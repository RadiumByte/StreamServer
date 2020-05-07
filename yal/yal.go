package yal

import (
	"os/exec"
)

// YoutubeClient represents data for connection to Youtube
type YoutubeClient struct {
}

// RunDevRTMP coverts video device stream to RTMP stream using FFMPEG
func (youtube *YoutubeClient) RunDevRTMP(deviceInput string, rtmpOutput string) {
	cmdKill := exec.Command("killall", "ffmpeg")
	cmdKill.Run()

	cmdKill = exec.Command("killall", "cvlc")
	cmdKill.Run()

	cmdFFMPEG := exec.Command("ffmpeg", "-f", "v4l2", "-i", deviceInput, "-ar", "44100", "-ac",
		"2", "-acodec", "pcm_s16le", "-f", "s16le", "-ac", "2", "-i", "/dev/zero", "-acodec", "aac", "-ab", "128k", "-strict", "experimental",
		"-aspect", "16:9", "-vcodec", "h264", "-preset", "veryfast", "-crf", "25", "-pix_fmt", "yuv420p", "-g", "60", "-vb", "820k", "-maxrate", "820k",
		"-bufsize", "820k", "-r", "30", "-f", "flv", rtmpOutput)
	cmdFFMPEG.Start()

	/*ffmpeg -thread_queue_size 512 -f v4l2 -i /dev/video2   -ar 44100 -ac
	2 -acodec pcm_s16le -f s16le -ac 2 -i /dev/zero -acodec aac -ab 128k -strict experimental
	-aspect 16:9 -vcodec h264 -preset veryfast -crf 25 -pix_fmt yuv420p -g 60 -vb 820k -maxrate 820k
	-bufsize 820k -profile:v baseline   -r 30 -f flv*/
}

// RunIPRTMP coverts input RTSP stream to RTMP stream using FFMPEG
func (youtube *YoutubeClient) RunIPRTMP(rtspInput string, cameraType int, rtmpOutput string) {
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

	cmdFFMPEG := exec.Command("ffmpeg", "-f", "lavfi", "-i", "anullsrc", "-rtsp_transport", transport,
		"-i", rtspInput, "-tune", "zerolatency", "-vcodec", "libx264", "-t", "12:00:00", "-pix_fmt",
		"+", "-c:v", "copy", "-c:a", "aac", "-strict", "experimental", "-f", "flv", rtmpOutput)
	cmdFFMPEG.Start()
}

//cvlc -vvv v4l2:///dev/video0 --sout '#transcode{vcodec=mp2v,vb=800,acodec=none}:rtp{sdp=rtsp://:8554/}'

// RunVLC coverts input /dev/video stream to RTSP stream using VLC
func (youtube *YoutubeClient) RunVLC(deviceInput string, rtspOutput string) {
	cmdKill := exec.Command("killall", "cvlc")
	cmdKill.Run()

	cmdKill = exec.Command("killall", "ffmpeg")
	cmdKill.Run()

	cmdVLC := exec.Command("cvlc", "-vvv", "v4l2://"+deviceInput, "--sout",
		"#transcode{vcodec=h264,vb=2000,acodec=none}:rtp{sdp="+rtspOutput+"}")
	cmdVLC.Start()
}

// NewYoutubeClient constructs object of YoutubeClient
func NewYoutubeClient() (*YoutubeClient, error) {
	res := &YoutubeClient{}

	return res, nil
}

package yal

// YoutubeClient represents data for connection to Youtube
type YoutubeClient struct {
}

// NewYoutubeClient constructs object of YoutubeClient
func NewYoutubeClient() (*YoutubeClient, error) {

	res := &YoutubeClient{}

	return res, nil
}

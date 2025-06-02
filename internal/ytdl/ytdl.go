package ytdl

import (
	"context"
	"strings"

	"github.com/kkdai/youtube/v2"
)

type YoutubeDL struct {
	client *youtube.Client
}

func NewYoutubeDL() *YoutubeDL {
	return &YoutubeDL{
		client: &youtube.Client{},
	}
}

func (dl *YoutubeDL) GetVideoInfo(ctx context.Context, videoID string) (*youtube.Video, string, error) {
	video, err := dl.client.GetVideoContext(ctx, videoID)
	if err != nil {
		return nil, "", err
	}
	formats := video.Formats.Select(func(format youtube.Format) bool {
		if strings.HasPrefix(format.MimeType, "video/") {
			return true
		}
		return false
	})
	if len(formats) == 0 {
		return nil, "", youtube.ErrNoFormat
	}
	format := formats[0]
	u, err := dl.client.GetStreamURLContext(ctx, video, &format)
	if err != nil {
		return nil, "", err
	}
	return video, u, nil
}

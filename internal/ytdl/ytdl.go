package ytdl

import (
	"context"
	_ "embed"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"strings"

	"github.com/kkdai/youtube/v2"
	"github.com/mengzhuo/cookiestxt"
)

type YoutubeDL struct {
	client *youtube.Client
}

// cookies.txt should be in the netscape cookie file format.
//
//go:embed cookies.txt
var initCookies string

func NewYoutubeDL() *YoutubeDL {
	jar, _ := cookiejar.New(nil)
	lines := strings.Split(initCookies, "\n")
	cookies := make([]*http.Cookie, 0, len(lines))
	for _, line := range lines {
		c, err := cookiestxt.ParseLine(line)
		if err != nil {
			continue // Skip invalid lines
		}
		if c == nil {
			continue // Skip empty cookies
		}
		cookies = append(cookies, c)
	}
	return &YoutubeDL{
		client: &youtube.Client{
			HTTPClient: &http.Client{
				Jar: jar,
			},
		},
	}
}

func (dl *YoutubeDL) GetVideoInfo(ctx context.Context, videoID string) (*youtube.Video, string, error) {
	video, err := dl.client.GetVideoContext(ctx,
		fmt.Sprintf("https://www.youtube.com/shorts/%s", videoID))
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

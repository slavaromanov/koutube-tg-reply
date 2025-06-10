package ytdl

import (
	"context"
	"crypto/tls"
	_ "embed"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"github.com/kkdai/youtube/v2"
	"github.com/mengzhuo/cookiestxt"
	"github.com/quic-go/quic-go/http3"
)

type YoutubeDL struct {
	client *youtube.Client
}

// cookies.txt should be in the netscape cookie file format.
//
//go:embed cookies.txt
var initCookies string

func newCookiesJar() (http.CookieJar, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(initCookies, "\n")
	cookies := make([]*http.Cookie, 0, len(lines))
	for _, line := range lines {
		c, err := cookiestxt.ParseLine(line)
		if err != nil || c == nil {
			continue // Skip invalid or empty cookies
		}
		cookies = append(cookies, c)
	}
	for _, host := range []string{
		"www.youtube.com",
		"youtube.com",
		"accounts.youtube.com",
	} {
		jar.SetCookies(&url.URL{
			Scheme: "https",
			Host:   host,
		}, cookies)
	}
	return jar, nil
}

func NewYoutubeDL() *YoutubeDL {
	jar, _ := newCookiesJar()
	return &YoutubeDL{
		client: &youtube.Client{
			HTTPClient: &http.Client{
				Jar: jar,
				Transport: &http3.Transport{
					TLSClientConfig: &tls.Config{
						InsecureSkipVerify: true,           // Disable certificate verification for testing
						NextProtos:         []string{"h3"}, // Indicate HTTP/3 support
					},
				},
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

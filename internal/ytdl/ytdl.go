package ytdl

import (
	"context"
	_ "embed"
	"net"
	"net/http"
	"net/url"
	"slices"
	"strings"
	"sync"
	"time"

	"github.com/kkdai/youtube/v2"
	"github.com/mengzhuo/cookiestxt"
)

type YoutubeDL struct {
	client *youtube.Client
}

var _ http.CookieJar = (*cookieJar)(nil)

type cookieJar struct {
	mu      sync.RWMutex
	cookies map[string][]*http.Cookie
}

func (c *cookieJar) SetCookies(u *url.URL, cookies []*http.Cookie) {
	for _, cookie := range cookies {
		if !slices.Contains(c.cookies[u.Host], cookie) {
			c.mu.Lock()
			c.cookies[u.Host] = append(c.cookies[u.String()], cookie)
			c.mu.Unlock()
		}
	}
}

func (c *cookieJar) Cookies(u *url.URL) []*http.Cookie {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if cookies, ok := c.cookies[strings.TrimPrefix(u.Host, "www")]; ok {
		return slices.Clone(cookies)
	}
	return nil
}

// cookies.txt should be in the netscape cookie file format.
//
//go:embed cookies.txt
var initCookies string

func newCookieJar() *cookieJar {
	jar := &cookieJar{
		cookies: make(map[string][]*http.Cookie),
	}
	lines := strings.Split(initCookies, "\n")
	for _, line := range lines {
		c, err := cookiestxt.ParseLine(line)
		if err != nil {
			continue // Skip invalid lines
		}
		if c == nil {
			continue // Skip empty cookies
		}
		jar.cookies[c.Domain] = append(jar.cookies[c.Domain], c)
	}
	return jar
}

func NewYoutubeDL() *YoutubeDL {
	return &YoutubeDL{
		client: &youtube.Client{
			HTTPClient: &http.Client{
				Transport: &http.Transport{
					DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
						dialer := &net.Dialer{
							Timeout:   30 * time.Second,
							KeepAlive: 30 * time.Second,
						}
						return dialer.DialContext(ctx, "tcp6", addr)
					},
				},
				// Jar: newCookieJar(),
			},
		},
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

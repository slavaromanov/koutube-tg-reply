package koutube_conv

import (
	"fmt"
	"regexp"
)

type VideoType uint8

const (
	Default VideoType = iota + 1
	Shorts
	Embed
	Channel
	Live
)

type VideoInfo struct {
	VideoID string
	Type    VideoType
}

type Converter struct {
	re        *regexp.Regexp
	groupKeys []string
}

func NewConverter() *Converter {
	re := regexp.
		MustCompile(`(?P<mobile>m\.)?youtu\.?be(\.com)?/(?P<browser>watch\?v=)?(?P<channel>c/|@)?(?P<stream>live/)?(?P<shorts>shorts/)?(?P<embed>embed/)?(?P<video_id>[0-9a-zA-Z_-]+).*(?P<playlist>list=[a-zA-Z0-9]+)?.*`)
	return &Converter{
		re:        re,
		groupKeys: re.SubexpNames(),
	}
}

func (c *Converter) ConvertVideoURL(s string) (bool, string) {
	if !c.re.MatchString(s) {
		return false, ""
	}
	info := c.getVideoInfo(s)
	switch info.Type {
	case Shorts:
		return true, fmt.Sprintf("https://glorytofight.ru/shorts/%s", info.VideoID)
	default:
		return false, ""
	}
}

func (c *Converter) getVideoInfo(s string) *VideoInfo {
	groups := c.mapFromSubmatch(c.re.FindStringSubmatch(s))
	switch {
	case groups["shorts"] != "":
		return &VideoInfo{
			VideoID: groups["video_id"],
			Type:    Shorts,
		}
	case groups["embed"] != "":
		return &VideoInfo{
			VideoID: groups["video_id"],
			Type:    Embed,
		}
	case groups["stream"] != "":
		return &VideoInfo{
			VideoID: groups["video_id"],
			Type:    Live,
		}
	case groups["channel"] != "":
		return &VideoInfo{
			VideoID: groups["video_id"],
			Type:    Channel,
		}
	default:
		return &VideoInfo{
			VideoID: groups["video_id"],
			Type:    Default,
		}
	}
}

func (c *Converter) mapFromSubmatch(submatch []string) map[string]string {
	m := make(map[string]string)
	for i, s := range submatch {
		m[c.groupKeys[i]] = s
	}
	return m
}

func (c *Converter) GetVideoID(s string) (string, error) {
	if !c.re.MatchString(s) {
		return "", fmt.Errorf("invalid video url")
	}
	info := c.getVideoInfo(s)
	return info.VideoID, nil
}

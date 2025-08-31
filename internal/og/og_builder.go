package og

import (
	"bytes"
	_ "embed"
	"html/template"

	"github.com/kkdai/youtube/v2"
)

type Builder struct {
	t *template.Template
}

//go:embed templ.gotmpl
var templ string

func NewBuilder() *Builder {
	return &Builder{
		t: template.Must(template.New("og_template").Parse(templ)),
	}
}

func (b *Builder) Build(video youtube.Video, streamURL string) (string, error) {
	buf := bytes.NewBuffer(nil)
	thumbnail := ""
	if len(video.Thumbnails) > 0 {
		thumbnail = video.Thumbnails[0].URL
	}
	err := b.t.Execute(buf, &MetaInfo{
		SiteName:       "𓂸ඞ PoopTube 𓂸ඞ",
		StreamURL:      streamURL,
		YoutubeVideoID: video.ID,
		Image:          thumbnail,
		Description:    video.Description,
		Author:         video.Author,
	})
	if err != nil {
		return "", err
	}
	defer buf.Reset()
	return buf.String(), nil
}

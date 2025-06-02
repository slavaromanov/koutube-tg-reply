package ytdl

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestYoutubeDL_GetVideoInfo(t *testing.T) {
	dl := NewYoutubeDL()
	ctx := context.Background()
	videoID := "e08I6tnSL84"
	info, err := dl.GetVideoInfo(ctx, videoID)
	ass := assert.New(t)
	ass.NoError(err)
	ass.NotEmpty(info, "Expected non-empty video info")
	t.Logf("Video Info: %s", info)
}

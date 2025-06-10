package ytdl

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestYoutubeDL_GetVideoInfo(t *testing.T) {
	dl := NewYoutubeDL()
	ctx := context.Background()
	videoID := "mehXi19as30"
	_, info, err := dl.GetVideoInfo(ctx, videoID)
	ass := assert.New(t)
	ass.NoError(err)
	ass.NotEmpty(info, "Expected non-empty video info")
	t.Logf("Video Info: %s", info)
}

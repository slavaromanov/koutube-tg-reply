package page

import (
	"context"

	"github.com/kkdai/youtube/v2"
)

type Builder interface {
	// Build generates the HTML page for the given video ID and stream URL.
	Build(video youtube.Video, videoID, streamURL string) (string, error)
}

type VideoStreamURLExtractor interface {
	// GetVideoInfo extracts the stream URL and metadata for the given video ID.
	GetVideoInfo(ctx context.Context, videoID string) (*youtube.Video, string, error)
}

type Service struct {
	builder  Builder
	streamer VideoStreamURLExtractor
}

func NewService(builder Builder, streamer VideoStreamURLExtractor) *Service {
	return &Service{
		builder:  builder,
		streamer: streamer,
	}
}

func (s *Service) BuildPage(ctx context.Context, videoID string) (string, error) {
	video, streamURL, err := s.streamer.GetVideoInfo(ctx, videoID)
	if err != nil {
		return "", err
	}

	html, err := s.builder.Build(*video, streamURL)
	if err != nil {
		return "", err
	}

	return html, nil
}

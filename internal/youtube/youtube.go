package youtube

import (
	"context"

	"time"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

type YouTubeWatcher struct {
	service *youtube.Service
}

func NewYouTubeWatcher(apiKey string) (*YouTubeWatcher, error) {
	ctx := context.Background()
	service, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}

	return &YouTubeWatcher{service: service}, nil
}

func (yw *YouTubeWatcher) CheckNewVideos(channelID string) ([]string, error) {
	call := yw.service.Search.List([]string{"snippet"}).
		ChannelId(channelID).
		Order("date").
		MaxResults(5).
		PublishedAfter(time.Now().Add(-24 * time.Hour).Format(time.RFC3339))

	response, err := call.Do()
	if err != nil {
		return nil, err
	}

	var videoIDs []string
	for _, item := range response.Items {
		if item.Id.Kind == "youtube#video" {
			videoIDs = append(videoIDs, item.Id.VideoId)
		}
	}

	return videoIDs, nil
}

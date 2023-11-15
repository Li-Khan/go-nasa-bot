package downloadVideo

import (
	"context"
	"fmt"
	"github.com/kkdai/youtube/v2"
	"io"
	"os"
	"time"
)

func Handle(videoURL string) error {
	client := youtube.Client{}
	video, err := client.GetVideo(videoURL)
	if err != nil {
		return err
	}
	if video == nil || video.Formats == nil || len(video.Formats) < 1 {
		return fmt.Errorf("video '%s' not found", videoURL)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()
	stream, _, err := client.GetStreamContext(ctx, video, &video.Formats[0])
	if err != nil {
		return err
	}

	file, err := os.Create("./video.mp4")
	if err != nil {
		return err
	}
	defer func() { _ = file.Close() }()

	_, err = io.Copy(file, stream)
	if err != nil {
		return err
	}
	return nil
}

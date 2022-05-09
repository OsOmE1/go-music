package youtube

import (
	"testing"
	"time"
)

func TestYouTubeDownloadParallel(t *testing.T) {
	ctx, err := NewContext(Youtube)
	if err != nil {
		t.Fatal(err)
	}
	res, err := ParseVideo("dQw4w9WgXcQ")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Video Result: \n"+
		"	Title: %s\n"+
		"	Length: %ds\n", res.Title, res.Length)
	formats, err := ctx.Formats(res)
	if err != nil {
		t.Fatal(err)
	}
	for _, format := range formats {
		t.Logf("Video Format: \n"+
			"	ContentLength: %d\n"+
			"	Duration: %ds\n"+
			"	Quality: %s\n"+
			"	Resolution %dx%d\n", format.ContentLength, format.DurationMS/1000, format.Quality, format.Width, format.Height)
	}
	// haven't found the optimal chunk size yet, but it seems to be around this 64k from looking at how Youtube handles their first data frame
	chunkSize := 64000
	t.Logf("downloading with chunkSize: %d", chunkSize)

	start := time.Now()
	if _, err := ctx.DownloadFormatParallel(formats[0], chunkSize); err != nil {
		t.Fatal(err)
	}
	t.Logf("download duration: %s", time.Now().Sub(start))
}

func TestYouTubeMusicDownloadParallel(t *testing.T) {
	ctx, err := NewContext(YoutubeMusic)
	if err != nil {
		t.Fatal(err)
	}
	res, err := ctx.ParseMusicTrack("dQw4w9WgXcQ")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Track Result: \n"+
		"	Title: %s\n"+
		"	Length: %ds\n", res.Title, res.Length)
	formats, err := ctx.Formats(res)
	if err != nil {
		t.Fatal(err)
	}
	for _, format := range formats {
		t.Logf("Track Format: \n"+
			"	ContentLength: %d\n"+
			"	Duration: %ds\n"+
			"	Quality: %s\n"+
			"	Bitrate %d\n", format.ContentLength, format.DurationMS/1000, format.AudioQuality, format.Bitrate)
	}
	// haven't found the optimal chunk size yet, but it seems to be around this 64k from looking at how Youtube handles their first data frame
	chunkSize := 64000
	t.Logf("downloading with chunkSize: %d", chunkSize)

	start := time.Now()
	if _, err := ctx.DownloadFormatParallel(formats[0], chunkSize); err != nil {
		t.Fatal(err)
	}
	t.Logf("download duration: %s", time.Now().Sub(start))
}

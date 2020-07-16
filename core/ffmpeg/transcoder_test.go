package ffmpeg

import (
	"testing"
)

func TestFFmpegCommand(t *testing.T) {
	transcoder := new(Transcoder)
	transcoder.ffmpegPath = "/fake/path/ffmpeg"
	transcoder.SetSegmentLength(4)
	transcoder.SetInput("fakecontent.flv")
	transcoder.SetOutputPath("fakeOutput")
	transcoder.SetHLSPlaylistLength(10)

	variant := HLSVariant{}
	variant.videoBitrate = "1200k"
	variant.isAudioPassthrough = true
	variant.encoderPreset = "veryfast"
	variant.SetVideoFramerate(30)

	transcoder.AddVariant(variant)

	cmd := transcoder.getString()

	expected := "cat fakecontent.flv | /fake/path/ffmpeg -hide_banner -i pipe:  -map v:0 -c:v:0 libx264 -b:v:0 1200k -map a:0 -c:a:0 copy -r 30 -g 60 -keyint_min 60 -preset veryfast  -var_stream_map \"v:0,a:0 \" -f hls -hls_time 4 -hls_list_size 10 -hls_delete_threshold 10 -hls_flags delete_segments+program_date_time+temp_file -tune zerolatency -sc_threshold 0 -master_pl_name stream.m3u8 -strftime 1 -hls_segment_filename fakeOutput/%v/stream-%s.ts -max_muxing_queue_size 400 fakeOutput/%v/stream.m3u8 2> transcoder.log"

	if cmd != expected {
		t.Errorf("ffmpeg command does not match expected.  Got %s, want: %s", cmd, expected)
	}
}

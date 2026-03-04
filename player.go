package main

import (
	"context"
	"fmt"
	"io"

	"github.com/ebitengine/oto/v3"
	"github.com/hajimehoshi/go-mp3"
)

func playAudio(ctx context.Context, source io.ReadCloser) error {
	defer source.Close()

	decoder, err := mp3.NewDecoder(source)
	if err != nil {
		return fmt.Errorf("decoding mp3: %w", err)
	}

	op := &oto.NewContextOptions{
		SampleRate:   decoder.SampleRate(),
		ChannelCount: 2,
		Format:       oto.FormatSignedInt16LE,
	}
	otoCtx, readyChan, err := oto.NewContext(op)
	if err != nil {
		return fmt.Errorf("creating audio context: %w", err)
	}
	<-readyChan

	player := otoCtx.NewPlayer(decoder)
	player.SetBufferSize(decoder.SampleRate() * 4 * 2) // ~2s at stereo 16-bit
	player.Play()

	<-ctx.Done()

	player.Pause()
	if err := otoCtx.Suspend(); err != nil {
		return fmt.Errorf("suspending audio: %w", err)
	}

	return nil
}

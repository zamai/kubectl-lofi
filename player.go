package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/ebitengine/oto/v3"
	"github.com/hajimehoshi/go-mp3"
)

// AudioEngine manages oto context (singleton) and swaps players on station switch.
type AudioEngine struct {
	mu      sync.Mutex
	otoCtx  *oto.Context
	player  *oto.Player
	stream  *StreamCloser
	cancel  context.CancelFunc
}

// StreamCloser bundles the HTTP body so we can close it on stop.
type StreamCloser struct {
	body interface{ Close() error }
}

func NewAudioEngine() (*AudioEngine, error) {
	// Create oto context with common defaults; the actual sample rate from
	// the first decoded stream will match 44100 for most lo-fi stations.
	op := &oto.NewContextOptions{
		SampleRate:   44100,
		ChannelCount: 2,
		Format:       oto.FormatSignedInt16LE,
	}
	ctx, readyChan, err := oto.NewContext(op)
	if err != nil {
		return nil, fmt.Errorf("creating audio context: %w", err)
	}
	<-readyChan

	return &AudioEngine{otoCtx: ctx}, nil
}

// Play stops current playback and starts a new stream. Non-blocking — returns
// once the stream is connected and decoding has begun. The onConnected callback
// is called (if non-nil) once audio starts playing.
func (ae *AudioEngine) Play(url string, onConnected func()) error {
	ae.Stop()

	ae.mu.Lock()
	defer ae.mu.Unlock()

	ctx, cancel := context.WithCancel(context.Background())
	ae.cancel = cancel

	body, err := openStream(ctx, url)
	if err != nil {
		cancel()
		return fmt.Errorf("opening stream: %w", err)
	}
	ae.stream = &StreamCloser{body: body}

	decoder, err := mp3.NewDecoder(body)
	if err != nil {
		body.Close()
		cancel()
		return fmt.Errorf("decoding mp3: %w", err)
	}

	player := ae.otoCtx.NewPlayer(decoder)
	player.SetBufferSize(decoder.SampleRate() * 4 * 2) // ~2s buffer
	player.Play()
	ae.player = player

	if onConnected != nil {
		onConnected()
	}

	return nil
}

// Stop pauses the current player and closes the stream.
func (ae *AudioEngine) Stop() {
	ae.mu.Lock()
	defer ae.mu.Unlock()

	if ae.cancel != nil {
		ae.cancel()
		ae.cancel = nil
	}
	if ae.player != nil {
		ae.player.Pause()
		ae.player = nil
	}
	if ae.stream != nil {
		ae.stream.body.Close()
		ae.stream = nil
	}
}

// Close suspends the oto context. Call on exit.
func (ae *AudioEngine) Close() {
	ae.Stop()
	ae.mu.Lock()
	defer ae.mu.Unlock()
	if ae.otoCtx != nil {
		_ = ae.otoCtx.Suspend()
	}
}

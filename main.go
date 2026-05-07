package main

import (
	"fmt"
	"os"
	"sync/atomic"
	"time"

	"golang.org/x/term"
)

func main() {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error entering raw mode: %v\n", err)
		os.Exit(1)
	}
	cleanup := func() {
		showCursor()
		term.Restore(int(os.Stdin.Fd()), oldState)
		fmt.Print("\r\n")
	}
	defer cleanup()

	hideCursor()

	engine, err := NewAudioEngine()
	if err != nil {
		cleanup()
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	defer engine.Close()

	selected := 0
	playing := 0 // which station is currently playing/connecting
	var connecting atomic.Bool
	connecting.Store(true)
	knobChar := byte('@')

	eqBars := randomEqBars(eqBarCount)

	// Channel for redraw signals from background goroutines.
	redrawCh := make(chan struct{}, 1)
	signalRedraw := func() {
		select {
		case redrawCh <- struct{}{}:
		default:
		}
	}

	renderFull(selected, true, eqBars, knobChar)

	// Start playing first station.
	go func() {
		if err := engine.Play(stations[selected].URL, func() {
			connecting.Store(false)
			signalRedraw()
		}); err != nil {
			connecting.Store(false)
			signalRedraw()
		}
	}()

	// Key input channel.
	keyCh := make(chan byte, 16)
	go func() {
		buf := make([]byte, 3)
		for {
			n, err := os.Stdin.Read(buf)
			if err != nil {
				return
			}
			for i := 0; i < n; i++ {
				keyCh <- buf[i]
			}
		}
	}()

	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	// Dial timer: after the user stops turning for 300ms, tune to the station.
	var dialTimer *time.Timer
	dialCh := make(chan struct{}, 1)

	startDialTimer := func() {
		if dialTimer != nil {
			dialTimer.Stop()
		}
		dialTimer = time.AfterFunc(300*time.Millisecond, func() {
			select {
			case dialCh <- struct{}{}:
			default:
			}
		})
	}

	// Rotate knob character on dial turn.
	knobIdx := 0
	rotateKnob := func() {
		knobIdx = (knobIdx + 1) % len(knobFrames)
		knobChar = knobFrames[knobIdx]
	}

	for {
		select {
		case b := <-keyCh:
			switch b {
			case 'q', 0x03: // q or Ctrl-C
				return
			case '\x1b': // escape sequence
				b2 := readKeyTimeout(keyCh, 50*time.Millisecond)
				if b2 != '[' {
					continue
				}
				b3 := readKeyTimeout(keyCh, 50*time.Millisecond)
				switch b3 {
				case 'A': // up
					if selected > 0 {
						selected--
						rotateKnob()
						renderFull(selected, selected != playing || connecting.Load(), eqBars, knobChar)
						startDialTimer()
					}
				case 'B': // down
					if selected < len(stations)-1 {
						selected++
						rotateKnob()
						renderFull(selected, selected != playing || connecting.Load(), eqBars, knobChar)
						startDialTimer()
					}
				}
			}

		case <-dialCh:
			// User stopped turning — tune to selected station.
			playing = selected
			connecting.Store(true)
			renderFull(selected, true, eqBars, knobChar)
			sel := selected
			go func() {
				if err := engine.Play(stations[sel].URL, func() {
					connecting.Store(false)
					signalRedraw()
				}); err != nil {
					connecting.Store(false)
					signalRedraw()
				}
			}()

		case <-redrawCh:
			renderFull(selected, connecting.Load(), eqBars, knobChar)

		case <-ticker.C:
			eqBars = randomEqBars(eqBarCount)
			renderEqualizer(eqBars)
		}
	}
}

func readKeyTimeout(ch <-chan byte, timeout time.Duration) byte {
	select {
	case b := <-ch:
		return b
	case <-time.After(timeout):
		return 0
	}
}

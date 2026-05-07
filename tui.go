package main

import (
	"fmt"
	"math/rand"
	"strings"
)

const (
	escClear      = "\033[2J"
	escHome       = "\033[H"
	escHideCursor = "\033[?25l"
	escShowCursor = "\033[?25h"
)

func clearScreen() { fmt.Print(escClear + escHome) }
func hideCursor()  { fmt.Print(escHideCursor) }
func showCursor()  { fmt.Print(escShowCursor) }
func moveTo(row, col int) { fmt.Printf("\033[%d;%dH", row, col) }

// rawPrint outputs text with \n replaced by \r\n for raw terminal mode.
func rawPrint(s string) {
	fmt.Print(strings.ReplaceAll(s, "\n", "\r\n"))
}

// renderFull draws the entire screen.
func renderFull(selected int, connecting bool, eqBars []int, knobChar byte) {
	clearScreen()
	renderBanner(selected, knobChar)
	renderStations(selected, connecting)
	renderEqualizer(eqBars)
	renderFooter()
}

func renderBanner(selected int, knobChar byte) {
	// Build the banner with dynamic station name on the "display" and knob char
	stName := stations[selected].Name
	// Pad or truncate to 16 chars for the display area
	if len(stName) > 16 {
		stName = stName[:16]
	}
	display1 := fmt.Sprintf("%-16s", stName)
	// Frequency-like line
	freq := fmt.Sprintf("%-16s", fmDial(selected))

	k := string(knobChar)

	rawPrint(fmt.Sprintf(`
  ╔══════════════════════════════════════════╗
  ║   🎵  k u b e c t l   l o - f i  🎵      ║
  ║                                          ║
  ║      ╭───────────────────────╮           ║
  ║      │  ♪ ♫ ♪ ♫ ♪ ♫ ♪ ♫ ♪    │           ║
  ║      │   beats to mass-      │           ║
  ║      │   scale clusters to   │           ║
  ║      ╰───────────────────────╯           ║
  ║                                          ║
  ║   ┌──────────────────┐  ╭───╮           ║
  ║   │  %s│  │   │           ║
  ║   │  %s│  │ %s │ << TUNE   ║
  ║   │  ≈≈≈≈≈≈≈≈≈≈≈≈≈≈≈≈│  │   │           ║
  ║   └──────────────────┘  ╰───╯           ║
  ╚══════════════════════════════════════════╝
`, display1, freq, k))
}

var fmFreqs = []string{
	"FM 87.5 MHz",
	"FM 94.3 MHz",
	"FM 101.1 MHz",
}

func fmDial(idx int) string {
	if idx < len(fmFreqs) {
		return fmFreqs[idx]
	}
	return fmt.Sprintf("FM %d", idx+1)
}

func renderStations(selected int, connecting bool) {
	rawPrint("\r\n")
	for i, s := range stations {
		marker := "  "
		if i == selected {
			marker = "⟩ "
		}
		line := fmt.Sprintf("    %s%s", marker, s.Name)
		if i == selected && connecting {
			line += "  \033[33m⟳ tuning...\033[0m"
		} else if i == selected {
			line += "  \033[32m♪ playing\033[0m"
		}
		rawPrint(line + "\r\n")
	}
	rawPrint("\r\n")
}

var eqBlocks = []rune("▁▂▃▄▅▆▇█")

func randomEqBars(n int) []int {
	bars := make([]int, n)
	for i := range bars {
		bars[i] = rand.Intn(len(eqBlocks))
	}
	return bars
}

const eqBarCount = 16
const eqRow = 24 // row where equalizer is drawn

func renderEqualizer(bars []int) {
	moveTo(eqRow, 1)
	fmt.Print("\033[2K")
	var b strings.Builder
	b.WriteString("    ")
	for _, h := range bars {
		if h < 0 {
			h = 0
		}
		if h >= len(eqBlocks) {
			h = len(eqBlocks) - 1
		}
		b.WriteRune(eqBlocks[h])
		b.WriteRune(' ')
	}
	fmt.Print(b.String())
}

func renderFooter() {
	moveTo(eqRow+2, 1)
	fmt.Print("  ↑/↓ tune stations • q quit")
}

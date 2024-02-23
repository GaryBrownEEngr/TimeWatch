package main

import (
	"fmt"
	_ "image/png"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type game struct {
	prevX              int
	prevY              int
	mouseJustMoved     bool
	lastTimeMouseMoved time.Time
	timeActive         time.Duration
	prevTime           time.Time
	updateNeeded       bool
	boxMsg             string
	titleMsg           string
}

const (
	width                  = 270
	height                 = 100
	maxInactiveTimeAllowed = 60 * 2
)

func (m *game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return width, height
}

func (m *game) Update() error {
	m.updateNeeded = true
	now := time.Now()

	mx, my := ebiten.CursorPosition()
	m.mouseJustMoved = mx != m.prevX || my != m.prevY
	if m.mouseJustMoved {
		m.lastTimeMouseMoved = now
	}

	timeSinceMoved := now.Sub(m.lastTimeMouseMoved).Seconds()
	if timeSinceMoved < maxInactiveTimeAllowed {

		m.timeActive += now.Sub(m.prevTime)
	}

	m.prevX = mx
	m.prevY = my

	m.prevTime = now

	//
	sec := m.timeActive.Seconds()
	min := sec / 60.0
	hour := min / 60.0
	m.boxMsg = fmt.Sprintf("Sec = %.0f\nMin =%.1f\nHour=%.2f", sec, min, hour)
	m.titleMsg = fmt.Sprintf("WR %.2f", hour)

	ebiten.SetWindowTitle(m.titleMsg)

	return nil
}

func (m *game) Draw(screen *ebiten.Image) {
	if !m.updateNeeded {
		return
	}
	m.updateNeeded = false
	screen.Clear()
	ebitenutil.DebugPrintAt(screen, m.boxMsg, 0, 0)
	// ebitenutil.DebugPrintAt(screen, fmt.Sprintf("x/y = %d/%d", m.prevX, m.prevY), 0, 65)
	// ebitenutil.DebugPrintAt(screen, fmt.Sprintf("FPS = %.0f\nTPS = %.0f", ebiten.ActualFPS(), ebiten.ActualTPS()), 0, 65)
}

func main() {
	ebiten.SetWindowSize(width, height)
	ebiten.SetTPS(1)
	ebiten.SetScreenClearedEveryFrame(false)
	ebiten.SetVsyncEnabled(false)
	ebiten.SetWindowTitle("WR 000000")

	n := &game{
		prevTime: time.Now(),
	}

	err := ebiten.RunGameWithOptions(n, &ebiten.RunGameOptions{
		//ScreenTransparent: true,
	})
	if err != nil {
		log.Fatal(err)
	}
}

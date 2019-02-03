package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"unicode/utf8"

	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/encoding"
)

var (
	defaultStyle               tcell.Style
	clockActive, clockInactive string
)

const (
	// ClockActiveDefault is default char used for active bit in clock
	ClockActiveDefault = "■"
	// ClockInactiveDefault is default char used for inactive bit in clock
	ClockInactiveDefault = "□"
)

// Util functions

func splitNum(n int) (int, int) {
	if n < 10 {
		return 0, n
	}
	return n / 10, n % 10
}

func getBin(n int) [4]int {
	var result [4]int
	next := 0
	for i := 0; i < 4; i++ {
		next = n / 2
		result[i] = n % 2
		n = next
	}
	return result
}

func main() {
	flag.StringVar(&clockActive, "o", ClockActiveDefault, "active bit char")
	flag.StringVar(&clockInactive, "z", ClockInactiveDefault, "inactive bit char")
	flag.Parse()

	encoding.Register()
	defaultStyle = tcell.StyleDefault.
		Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)

	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("Error when creating screen: %s", err.Error())
		os.Exit(1)
	}

	err = s.Init()
	defer s.Fini()
	if err != nil {
		log.Fatalf("Error initializing screen: %s", err.Error())
		os.Exit(1)
	}

	s.HideCursor()
	defer s.ShowCursor(1, 1)
	s.Clear()

	on, _ := utf8.DecodeRuneInString(clockActive)
	off, _ := utf8.DecodeLastRuneInString(clockInactive)

	cfg := &ClockWidgetConfig{
		x:          0,
		y:          0,
		padX:       2,
		padY:       2,
		sectionPad: 2,
		bitOn:      on,
		bitOff:     off,
	}

	clock := newClockWidget(cfg)
	clock.CenterPos(s)
	clock.Update()
	clock.Draw(s, &defaultStyle)
	s.Show()

	events := make(chan tcell.Event)
	signals := make(chan os.Signal, 2)
	signal.Notify(signals, syscall.SIGTERM, os.Interrupt)

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	go func() {
		for {
			events <- s.PollEvent()
		}
	}()

loop:
	for {
		select {
		case <-signals:
			break loop
		case ev := <-events:
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyEscape, tcell.KeyCtrlC:
					break loop
				}
			case *tcell.EventResize:
				s.Clear()
				clock.CenterPos(s)
				clock.Update()
				clock.Draw(s, &defaultStyle)
				s.Show()
			}
		case <-ticker.C:
			clock.Update()
			clock.Draw(s, &defaultStyle)
			s.Show()
		}
	}

}

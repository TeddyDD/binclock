package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"strings"
	"syscall"
	"time"
	"unicode/utf8"

	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/encoding"
)

var (
	defaultStyle               tcell.Style
	clockActive, clockInactive string

	reverseStyle bool
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

// https://github.com/golang/go/issues/20455
func fixTimezone() {
	if runtime.GOOS != "android" {
		return
	}
	out, err := exec.Command("/system/bin/getprop", "persist.sys.timezone").Output()
	if err != nil {
		return
	}
	z, err := time.LoadLocation(strings.TrimSpace(string(out)))
	if err != nil {
		return
	}
	time.Local = z
}

func main() {
	fixTimezone()
	flag.StringVar(&clockActive, "o", ClockActiveDefault, "active bit char")
	flag.StringVar(&clockInactive, "z", ClockInactiveDefault, "inactive bit char")
	flag.BoolVar(&reverseStyle, "r", false, "reverse colors for clock bits")
	flag.Parse()

	encoding.Register()
	bg, fg := tcell.ColorBlack, tcell.ColorWhite
	if reverseStyle {
		bg, fg = tcell.ColorWhite, tcell.ColorBlack
	}
	defaultStyle = tcell.StyleDefault.
		Background(bg).Foreground(fg)

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
	off, _ := utf8.DecodeRuneInString(clockInactive)

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
				if ev.Rune() == 'q' {
					break loop
				}
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

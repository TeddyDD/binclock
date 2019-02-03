package main

import (
	"fmt"
	"github.com/gdamore/tcell"
	"strings"
	"time"
)

// Clock keeps current time and 2d array of bits
// used to display binary clock
type Clock struct {
	now     time.Time
	Display [6][4]int
}

func (c Clock) String() string {
	s := &strings.Builder{}
	fmt.Fprintf(s, "Clock: %s\n", c.now)
	for y := 3; y >= 0; y-- {
		for x := 0; x < 6; x++ {
			fmt.Fprintf(s, " %d ", c.Display[x][y])
		}
		fmt.Fprintf(s, "\n")
	}
	return s.String()
}

// Update sets internal clock time to tim.Now() and updates bit array
func (c *Clock) Update() {
	c.now = time.Now()
	c.updateDisplaySection(0, c.now.Hour())
	c.updateDisplaySection(1, c.now.Minute())
	c.updateDisplaySection(2, c.now.Second())
}

func (c *Clock) updateDisplaySection(section, number int) {
	a, b := splitNum(number)
	s := section * 2
	c.Display[s] = getBin(a)
	c.Display[s+1] = getBin(b)
}

type ClockWidgetConfig struct {
	x, y, padX, padY, sectionPad int
	bitOn                        rune
	bitOff                       rune
}

// ClockWidget is struct used to display Clock on tcell screen
type ClockWidget struct {
	*ClockWidgetConfig
	*Clock
}

// Draw clock on provided tcell screen with given style
func (c ClockWidget) Draw(s tcell.Screen, style *tcell.Style) {
	cx, cy := c.x, c.y
	for y := 3; y >= 0; y-- {
		for x := 0; x < 6; x++ {
			num := c.Display[x][y]
			if num == 0 {
				s.SetContent(cx, cy, ClockInactiveDefault, nil, *style)
			} else {
				s.SetContent(cx, cy, ClockActiveDefault, nil, *style)
			}
			cx += c.padX
			if x == 1 || x == 3 {
				cx += c.sectionPad
			}
		}
		cx = c.x
		cy += c.padY
	}
}

func (c ClockWidget) size() (int, int) {
	return (c.padX * 4) + (c.sectionPad * 2), (c.padY * 4)
}

// CenterPos changes x and y values of ClockWidget to
// keep ClockWidget centered. Call this after every screen resize event
func (c *ClockWidget) CenterPos(s tcell.Screen) {
	screenW, screenH := s.Size()
	wW, wH := c.size()
	c.x = (screenW / 2) - (wW / 2)
	c.y = (screenH / 2) - (wH / 2)
}

func newClockWidget(config *ClockWidgetConfig) *ClockWidget {
	c := &ClockWidget{
		config,
		&Clock{},
	}
	return c
}

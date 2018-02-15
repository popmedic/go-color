package tty

import (
	"fmt"

	"github.com/popmedic/go-color/colorize"
)

// TTY is an int holding number escape value for a color on the terminal.
type TTY int

// Font Types
const (
	reset TTY = iota
	bold
	faint
	italic
	underline
	blinkSlow
	blinkRapid
	reverseVideo
	concealed
	crossedOut
)

func Reset() colorize.IColorize        { return newTTYColorize(reset) }
func Bold() colorize.IColorize         { return newTTYColorize(bold) }
func Faint() colorize.IColorize        { return newTTYColorize(faint) }
func Italic() colorize.IColorize       { return newTTYColorize(italic) }
func Underline() colorize.IColorize    { return newTTYColorize(underline) }
func BlinkSlow() colorize.IColorize    { return newTTYColorize(blinkSlow) }
func BlinkRapid() colorize.IColorize   { return newTTYColorize(blinkRapid) }
func ReverseVideo() colorize.IColorize { return newTTYColorize(reverseVideo) }
func Concealed() colorize.IColorize    { return newTTYColorize(concealed) }
func CrossedOut() colorize.IColorize   { return newTTYColorize(crossedOut) }

// Foreground Colors
const (
	fgBlack TTY = iota + 30
	fgRed
	fgGreen
	fgYellow
	fgBlue
	fgMagenta
	fgCyan
	fgWhite
)

func FgBlack() colorize.IColorize   { return newTTYColorize(fgBlack) }
func FgRed() colorize.IColorize     { return newTTYColorize(fgRed) }
func FgGreen() colorize.IColorize   { return newTTYColorize(fgGreen) }
func FgYellow() colorize.IColorize  { return newTTYColorize(fgYellow) }
func FgBlue() colorize.IColorize    { return newTTYColorize(fgBlue) }
func FgMagenta() colorize.IColorize { return newTTYColorize(fgMagenta) }
func FgCyan() colorize.IColorize    { return newTTYColorize(fgCyan) }
func FgWhite() colorize.IColorize   { return newTTYColorize(fgWhite) }

// Foreground High Colors
const (
	fgHiBlack TTY = iota + 90
	fgHiRed
	fgHiGreen
	fgHiYellow
	fgHiBlue
	fgHiMagenta
	fgHiCyan
	fgHiWhite
)

func FgHiBlack() colorize.IColorize   { return newTTYColorize(fgHiBlack) }
func FgHiRed() colorize.IColorize     { return newTTYColorize(fgHiRed) }
func FgHiGreen() colorize.IColorize   { return newTTYColorize(fgHiGreen) }
func FgHiYellow() colorize.IColorize  { return newTTYColorize(fgHiYellow) }
func FgHiBlue() colorize.IColorize    { return newTTYColorize(fgHiBlue) }
func FgHiMagenta() colorize.IColorize { return newTTYColorize(fgHiMagenta) }
func FgHiCyan() colorize.IColorize    { return newTTYColorize(fgHiCyan) }
func FgHiWhite() colorize.IColorize   { return newTTYColorize(fgHiWhite) }

// Background Colors
const (
	bgBlack TTY = iota + 40
	bgRed
	bgGreen
	bgYellow
	bgBlue
	bgMagenta
	bgCyan
	bgWhite
)

func BgBlack() colorize.IColorize   { return newTTYColorize(bgBlack) }
func BgRed() colorize.IColorize     { return newTTYColorize(bgRed) }
func BgGreen() colorize.IColorize   { return newTTYColorize(bgGreen) }
func BgYellow() colorize.IColorize  { return newTTYColorize(bgYellow) }
func BgBlue() colorize.IColorize    { return newTTYColorize(bgBlue) }
func BgMagenta() colorize.IColorize { return newTTYColorize(bgMagenta) }
func BgCyan() colorize.IColorize    { return newTTYColorize(bgCyan) }
func BgWhite() colorize.IColorize   { return newTTYColorize(bgWhite) }

// Background High Colors
const (
	bgHiBlack TTY = iota + 100
	bgHiRed
	bgHiGreen
	bgHiYellow
	bgHiBlue
	bgHiMagenta
	bgHiCyan
	bgHiWhite
)

func BgHiBlack() colorize.IColorize   { return newTTYColorize(bgHiBlack) }
func BgHiRed() colorize.IColorize     { return newTTYColorize(bgHiRed) }
func BgHiGreen() colorize.IColorize   { return newTTYColorize(bgHiGreen) }
func BgHiYellow() colorize.IColorize  { return newTTYColorize(bgHiYellow) }
func BgHiBlue() colorize.IColorize    { return newTTYColorize(bgHiBlue) }
func BgHiMagenta() colorize.IColorize { return newTTYColorize(bgHiMagenta) }
func BgHiCyan() colorize.IColorize    { return newTTYColorize(bgHiCyan) }
func BgHiWhite() colorize.IColorize   { return newTTYColorize(bgHiWhite) }

const (
	unset  = "\x1b[0m"
	format = "\x1b[%d;1m"
)

// String returns the escape for a color as a string.
func (c TTY) String() string {
	if c == reset {
		return unset
	}
	return fmt.Sprintf(format, c)
}

func newTTYColorize(start TTY, color ...colorize.IColorize) colorize.IColorize {
	return colorize.NewColorize(start.String(), reset.String(), color...)
}

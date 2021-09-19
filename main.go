package main

import (
	"cart/game"
	"cart/palettes"
	"cart/w4"
	"strconv"
)

var ui *UI

//go:export start
func start() {
	seed := uint32(478194671)

	ui = newUI(game.New(seed), w4.GAMEPAD1, palettes.Platinum)
}

//go:export update
func update() {
	ui.check()
	ui.show()
	ui.input()
}

type UI struct {
	game.Board

	frame uint
	old   uint8
	pad   *uint8
}

func newUI(board game.Board, pad *uint8, palette [4]uint32) *UI {
	return &UI{Board: board, pad: pad}

	ui.changePalette(palette)

	return ui
}

func (ui *UI) check() {
	ui.frame++
}

func (ui *UI) input() {
	switch {
	case ui.btn1():
		log("BTN1")
		ui.Input(game.KeyBTN1)
		ui.randomPalette()
	case ui.btn2():
		log("BTN2")
		ui.Input(game.KeyBTN2)
	case ui.up():
		log("🡱")
		ui.Input(game.KeyUp)
	case ui.down():
		log("🡳")
		ui.Input(game.KeyDown)
	case ui.right():
		log("🡲")
		ui.Input(game.KeyRight)
	case ui.left():
		log("🡰")
		ui.Input(game.KeyLeft)
	}

	ui.old = *ui.pad
}

func (ui *UI) show() {
	dotbg(0, 0, 160, 27, 3, 0x3, 0x4)

	// Pts border
	color(0x32)
	rect(0, 6, 84, 12)

	// Bottom background
	color(0x4)
	rect(0, 157, 160, 3)

	color(0x24)
	text("Total"+leftpad(itoa(ui.Total()), " ", 5), 2, 8)

	color(0x41)
	text(" -+-+-+- ", 88, 0)
	color(0x42)
	text(" 2|0|4|8 ", 88, 8)
	color(0x43)
	text(" -+-+-+- ", 88, 16)

	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			if v := ui.Get(r, c); v > 0 {
				showTile(r, c, v)
			}
		}
	}

	color(0x3)

	// Horizontal lines on the board
	line(0, 28, 160, 28)
	line(0, 60, 160, 60)
	line(0, 92, 160, 92)
	line(0, 124, 160, 124)
	line(0, 156, 160, 156)

	color(0x2)
	line(0, 27, 160, 27)
	line(0, 59, 160, 59)
	line(0, 91, 160, 91)
	line(0, 123, 160, 123)
	line(0, 155, 160, 155)

	// Vertical lines on the board
	line(39, 28, 39, 155)
	line(80, 28, 80, 155)
	line(120, 28, 120, 155)

	if ui.frame < 124 {
		var (
			s  string
			fg uint16
			bg uint16
			a  int
			h  uint
		)

		switch {
		case ui.frame < 6 || ui.frame > 118:
			s = "L   E   T   S      G   O."
			fg, bg, a, h = 0x2, 0x21, 120, 10
		case ui.frame < 12 || ui.frame > 112:
			s = "L  E  T  S    G  O!"
			fg, bg, a, h = 0x3, 0x21, 115, 15
		case ui.frame < 18 || ui.frame > 106:
			s = "L E T S  G O!"
			fg, bg, a, h = 0x3, 0x32, 110, 20
		case ui.frame < 24 || ui.frame > 100:
			s = "LETS GO!"
			fg, bg, a, h = 0x4, 0x23, 105, 25
		default:
			s = "LETS GO!"
			fg, bg, a, h = 0x1, 0x34, 100, 30
		}

		color(bg)
		rect(8, a, 144, h)
		color(fg)
		text(s, 80-(len(s)*8/2), a+(int(h)/2)-5)
	}
}

func (ui *UI) up() bool {
	return *ui.pad&w4.BUTTON_UP != 0 && ui.old&w4.BUTTON_UP == 0
}

func (ui *UI) down() bool {
	return *ui.pad&w4.BUTTON_DOWN != 0 && ui.old&w4.BUTTON_DOWN == 0
}

func (ui *UI) right() bool {
	return *ui.pad&w4.BUTTON_RIGHT != 0 && ui.old&w4.BUTTON_RIGHT == 0
}

func (ui *UI) left() bool {
	return *ui.pad&w4.BUTTON_LEFT != 0 && ui.old&w4.BUTTON_LEFT == 0
}

func (ui *UI) btn1() bool {
	return *ui.pad&w4.BUTTON_1 != 0 && ui.old&w4.BUTTON_1 == 0
}

func (ui *UI) btn2() bool {
	return *ui.pad&w4.BUTTON_2 != 0 && ui.old&w4.BUTTON_2 == 0
}

func (ui *UI) log(s string) {
	log(leftpad(utoa(ui.frame), " ", 5) + " - " + s)
}

func (ui *UI) changePalette(p [4]uint32) {
	w4.PALETTE[0] = p[0]
	w4.PALETTE[1] = p[1]
	w4.PALETTE[2] = p[2]
	w4.PALETTE[3] = p[3]
}

func (ui *UI) randomPalette() {
	i := int(ui.frame) % len(palettes.All)

	ui.changePalette(palettes.All[i])
}

func showTile(col, row, val int) {
	x, y := 2+(col*41), 32+row*32
	w, h := uint(35), uint(27)
	s := 8

	tileShadow(val)
	rect(x-1, y-1, w, h)

	set(x, y, 0x43)

	color(0x1)
	line(x-1, y, x, y-1)
	line(x+int(w-3), y-1, x+int(w-2), y)

	set(x-1, y-1, 0x1)               // Top-Left
	set(x+int(w-2), y-1, 0x1)        // Top-Right
	set(x-1, y+int(h)-2, 0x1)        // Bottom-Left
	set(x+int(w-2), y+int(h)-2, 0x1) // Bottom-Right

	tileColor(val)
	text("    ", x, y)
	text(leftpad(itoa(val), " ", 4), x, y+s)
	text("    ", x, y+s*2)
}

func tileColor(val int) {
	switch val {
	case 2:
		color(0x21)
	case 4:
		color(0x32)
	case 8:
		color(0x24)
	case 16:
		color(0x32)
	case 32:
		color(0x34)
	case 64:
		color(0x32)
	case 128:
		color(0x40)
	case 256:
		color(0x42)
	case 512:
		color(0x43)
	case 1024:
		color(0x34)
	case 2048:
		color(0x24)
	}
}

func tileShadow(val int) {
	switch val {
	case 2:
		color(0x43)
	case 4:
		color(0x43)
	case 8:
		color(0x42)
	case 16:
		color(0x32)
	case 32:
		color(0x34)
	case 64:
		color(0x32)
	case 128:
		color(0x43)
	case 256:
		color(0x42)
	case 512:
		color(0x24)
	case 1024:
		color(0x44)
	case 2048:
		color(0x44)
	}
}

func dotbg(x1, y1, w, h, s int, dc, bg uint16) {
	for x := x1; x < w; x++ {
		for y := y1; y < h; y++ {
			if x%s == 0 && y%s == 0 {
				set(x, y, dc)
			} else {
				set(x, y, bg)
			}
		}
	}
}

func set(x, y int, c uint16) {
	color(c)
	line(x, y, x, y)
}

func color(c uint16) {
	*w4.DRAW_COLORS = c
}

func line(x1, y1, x2, y2 int) {
	w4.Line(x1, y1, x2, y2)
}

func rect(x, y int, w, h uint) {
	w4.Rect(x, y, w, h)
}

func text(s string, x, y int) {
	w4.Text(s, x, y)
}

func log(s string) {
	w4.Trace(s)
}

func leftpad(s, c string, w int) string {
	n := w - len(s)

	if n <= 0 {
		return s
	}

	var p string

	for i := 0; i < n; i++ {
		p += c
	}

	return p + s
}

func itoa(i int) string {
	return strconv.Itoa(i)
}

func utoa(u uint) string {
	return strconv.FormatUint(uint64(u), 10)
}

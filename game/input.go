package game

type Key int

const (
	KeyUp Key = iota
	KeyDown
	KeyRight
	KeyLeft
	KeyBTN1
	KeyBTN2
)

func (b *board) Input(key Key) {
	switch key {
	case KeyUp:
		b.up()
		b.Add()
	case KeyDown:
		b.down()
		b.Add()
	case KeyRight:
		b.right()
		b.Add()
	case KeyLeft:
		b.left()
		b.Add()
	case KeyBTN1:
		// Do something
	case KeyBTN2:
		b.Restart()
	}
}

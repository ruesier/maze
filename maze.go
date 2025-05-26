package maze

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const TilePxSize = 50

// maze is 30 x 15 tiles
const Width = 30
const Widthpx = Width * TilePxSize
const Height = 15
const Heightpx = Height * TilePxSize

type State byte

const (
	START State = iota
	PLAY
	FINISH
)

type Maze struct {
	State
	m Layout
}

func (maze *Maze) Update() error {
	return nil
}

func (maze *Maze) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello World")
}

func (maze *Maze) Layout(_, _ int) (int, int) {
	return Widthpx, Heightpx
}

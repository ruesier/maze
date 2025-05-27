package maze

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
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
	m       Layout
	current struct {
		row int
		col int
	}
}

func NewMaze() *Maze {
	return &Maze{
		m: NewLayout(),
		current: struct {
			row int
			col int
		}{
			row: 2,
			col: 4,
		},
	}
}

func (maze *Maze) Update() error {
	return nil
}

func DrawTile(screen *ebiten.Image, row, col int, color color.Color) {
	cornerX := float32(col * TilePxSize)
	cornerY := float32(row * TilePxSize)
	vector.DrawFilledRect(screen, cornerX, cornerY, float32(TilePxSize), float32(TilePxSize), color, false)
}

func (maze *Maze) drawPlayer(screen *ebiten.Image) {
	CenterX := float32(maze.current.col*TilePxSize + TilePxSize/2)
	CenterY := float32(maze.current.row*TilePxSize + TilePxSize/2)
	vector.DrawFilledCircle(screen, CenterX, CenterY, 20, color.RGBA{
		R: 255,
		A: 255,
	}, true)
}

func (maze *Maze) drawFinish(screen *ebiten.Image) {
	CornerX := float32((Width-1)*TilePxSize + 10)
	CornerY := float32((Height-1)*TilePxSize + 10)
	vector.DrawFilledRect(
		screen,
		CornerX,
		CornerY,
		float32(TilePxSize-20),
		float32(TilePxSize-20),
		color.RGBA{G: 255, A: 255},
		false,
	)
}

func (maze *Maze) Draw(screen *ebiten.Image) {
	for r, row := range maze.m {
		for c, tile := range row {
			if tile {
				DrawTile(screen, r, c, color.White)
			} else {
				DrawTile(screen, r, c, color.Black)
			}
		}
	}
	maze.drawFinish(screen)
	maze.drawPlayer(screen)
}

func (maze *Maze) Layout(_, _ int) (int, int) {
	return Widthpx, Heightpx
}

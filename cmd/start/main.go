package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ruesier/maze"
)

func main() {
	ebiten.SetWindowSize(maze.Widthpx, maze.Heightpx)
	ebiten.SetWindowTitle("Maze")

	game := &maze.Maze{}
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}

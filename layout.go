package maze

import (
	"encoding/binary"
	"math/rand/v2"
	"time"
)

type Layout [][]bool // layout[row][col] = open

func NewLayout() Layout {
	nano := time.Now().UnixNano()
	var seed [32]byte
	seedSlice := seed[:0]
	seedSlice = binary.BigEndian.AppendUint64(seedSlice, uint64(nano))
	nano = time.Now().UnixNano()
	seedSlice = binary.BigEndian.AppendUint64(seedSlice, uint64(nano))
	nano = time.Now().UnixNano()
	seedSlice = binary.LittleEndian.AppendUint64(seedSlice, uint64(nano))
	nano = time.Now().UnixNano()
	binary.LittleEndian.AppendUint64(seedSlice, uint64(nano))

	gen := rand.New(rand.NewChaCha8(seed))

	out := make(Layout, Height)
	for row := range out {
		Row := make([]bool, 0, Width)
		for i := 0; len(Row) < Width && i*64 < Width; i++ {
			num := gen.Uint64()
			mask := uint64(1)
			for j := 0; len(Row) < Width && j < 64; j++ {
				Row = append(Row, num&mask > 0)
				mask = mask << 1
			}
		}
		out[row] = Row
	}

	solution := newSolution(gen)
	for r, row := range out {
		for c := range row {
			row[c] = solution[r][c] || row[c]
		}
	}

	return out
}

func newSolution(gen *rand.Rand) Layout {
	solution := make(Layout, 0, Height)
	for range Height {
		solution = append(solution, make([]bool, Width))
	}
	return randomWalk(solution, 0, 0, gen)
}

func randomWalk(visited Layout, row, col int, gen *rand.Rand) Layout {
	visited[row][col] = true
	if row == Height-1 && col == Width-1 {
		solution := make(Layout, 0, Height)
		for range Height {
			solution = append(solution, make([]bool, Width))
		}
		solution[row][col] = true
		return solution
	}
	var possibleNext [][2]int
	if row > 0 && !visited[row-1][col] {
		possibleNext = append(possibleNext, [2]int{row-1, col})
	}
	if row < Height - 1 && !visited[row+1][col] {
		possibleNext = append(possibleNext, [2]int{row+1, col})
	}
	if col > 0 && !visited[row][col-1] {
		possibleNext = append(possibleNext, [2]int{row, col-1})
	}
	if col < Width-1 && !visited[row][col+1] {
		possibleNext = append(possibleNext, [2]int{row, col+1})
	}
	gen.Shuffle(len(possibleNext), func(i, j int) {
		possibleNext[i], possibleNext[j] = possibleNext[j], possibleNext[i]
	})
	for _, next := range possibleNext {
		done := randomWalk(visited, next[0], next[1], gen)
		if done != nil{
			done[row][col] = true
			return done
		}
	}
	return nil
}
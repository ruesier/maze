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

	return out
}

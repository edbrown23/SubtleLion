package subtlelion

import (
    "image"
)

type chunk struct {
    cells [][]byte
}

type Point image.Point

type world struct {
    Chunks map[Point]chunk
}

func newChunk() *chunk {
    c := chunk{cells: make([][]byte, 8)}
    for i := range c.cells {
        c.cells[i] = make([]byte, 8)
    }
    return &c
}


func NewWorld() *world {
    w := world{Chunks: make(map[Point]chunk)}
    for y := -7; y < 8; y++ {
        for x := -7; x < 8; x++ {
            c := newChunk()
            p := Point{x, y}
            w.Chunks[p] = *c
        }
    }
    return &w
}


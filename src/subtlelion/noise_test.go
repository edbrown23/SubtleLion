package subtlelion

import (
    "fmt"
    "testing"
)

func TestPerlinNoise(t *testing.T) {
    for y := -50; y < 50; y++ {
        for x := -50; x < 50; x++ {
            noise1 := PerlinNoise(float64(x), float64(y))
            noise2 := PerlinNoise(float64(x), float64(y))
            if noise1 != noise2 {
                fmt.Printf("Noise1 %f Noise2 %f\n", noise1, noise2)
                t.FailNow()
            }
        }
    }
}

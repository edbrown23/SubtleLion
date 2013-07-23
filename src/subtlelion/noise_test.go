package subtlelion

import (
    "fmt"
    "image"
    "io/ioutil"
    "testing"
)

func TestPerlinNoise(t *testing.T) {
    boundsRect := image.Rectangle{Min: image.Point{0, 0},
                                  Max: image.Point{50, 50}}
    image := image.NewGray(boundsRect)
    bounds := image.Bounds()
    for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
        for x := bounds.Min.X; x < bounds.Max.X; x++ {
            noise := PerlinNoise(float64(x), float64(y))
            fmt.Println(noise)
            image.Pix[y * image.Stride + x] = uint8(noise)
        }
    }
    ioutil.WriteFile("output.jpeg", image.Pix, 0644)
    t.Fail()
}

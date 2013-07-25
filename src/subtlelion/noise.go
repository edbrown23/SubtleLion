package subtlelion

import (
    "math"
    "math/rand"
)


func cosineInterpolate(a, b, x float64) float64 {
    ft := x * math.Pi
    f := (1 - math.Cos(ft)) * 0.5

    return a * (1 - f) + b * f
}

func randomNoise(x, y int) float64 {
    return rand.Float64()
}

func smoothNoise(x, y float64) float64 {
    corners := (randomNoise(int(x - 1), int(y - 1)) + randomNoise(int(x + 1), int(y - 1)) + randomNoise(int(x - 1), int(y + 1)) + randomNoise(int(x + 1), int(y + 1))) / 16
    sides := (randomNoise(int(x - 1), int(y)) + randomNoise(int(x + 1), int(y)) + randomNoise(int(x), int(y - 1)) + randomNoise(int(x), int(y + 1))) / 8
    center := randomNoise(int(x), int(y)) / 4
    return corners + sides + center
}

func interpolatedNoise(x, y float64) float64 {
    intX := int(x)
    fracX := x - float64(intX)

    intY := int(y)
    fracY := y - float64(intY)

    v1 := smoothNoise(float64(intX), float64(intY))
    v2 := smoothNoise(float64(intX + 1), float64(intY))
    v3 := smoothNoise(float64(intX), float64(intY + 1))
    v4 := smoothNoise(float64(intX + 1), float64(intY + 1))

    i1 := cosineInterpolate(v1, v2, fracX)
    i2 := cosineInterpolate(v3, v4, fracX)

    return cosineInterpolate(i1, i2, fracY)
}

func PerlinNoise(x, y float64) float64 {
    total := 0.0
    persistence := 0.7
    octaves := 8

    for i := 0; i < octaves; i++ {
        frequency := math.Pow(2, float64(i))
        amplitude := math.Pow(persistence, float64(i))

        total += interpolatedNoise(x * frequency, y * frequency) * amplitude
    }

    return total / 1.7
}

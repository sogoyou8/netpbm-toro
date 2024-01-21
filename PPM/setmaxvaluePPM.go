package PPM

import "fmt"

// SetMaxValue sets the maximum value of the ppm image pixels.
func (ppm *PPM) SetMaxValue(maxValue uint8) {
	oldMax := ppm.max
	ppm.max = uint(maxValue)

	for y := 0; y < ppm.height; y++ {
		for x := 0; x < ppm.width; x++ {
			ppm.data[y][x].R = uint8(float64(ppm.data[y][x].R) * float64(ppm.max) / float64(oldMax))
			ppm.data[y][x].G = uint8(float64(ppm.data[y][x].G) * float64(ppm.max) / float64(oldMax))
			ppm.data[y][x].B = uint8(float64(ppm.data[y][x].B) * float64(ppm.max) / float64(oldMax))

		}
		fmt.Println(ppm.max)
	}
}
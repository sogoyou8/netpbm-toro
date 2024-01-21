package PPM

import "fmt"

func (ppm *PPM) Invert() {
	for x := 0; x < ppm.width; x++ {
		for y := 0; y < ppm.height; y++ {
			if ppm.magicNumber == "P2" {
				ppm.data[x][y].R = uint8(ppm.max) - ppm.data[x][y].R
				ppm.data[x][y].G = uint8(ppm.max) - ppm.data[x][y].G
				ppm.data[x][y].B = uint8(ppm.max) - ppm.data[x][y].B

			}
		}
		fmt.Println()
	}
}
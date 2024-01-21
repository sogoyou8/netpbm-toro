package PPM

import "fmt"

func (ppm *PPM) Flop() {
	for x := 0; x < ppm.width; x++ {
		for y := ppm.height - 1; y >= 0; y-- {
			if ppm.magicNumber == "P5" {
				fmt.Print(uint8(ppm.max)-ppm.data[x][y].R, " ")
				fmt.Print(uint8(ppm.max)-ppm.data[x][y].G, " ")
				fmt.Print(uint8(ppm.max)-ppm.data[x][y].B, " ")
				ppm.data[x][y].R = uint8(ppm.max) - ppm.data[x][y].R
				ppm.data[x][y].G = uint8(ppm.max) - ppm.data[x][y].G
				ppm.data[x][y].B = uint8(ppm.max) - ppm.data[x][y].B

			}
		}
		fmt.Println()
	}
}
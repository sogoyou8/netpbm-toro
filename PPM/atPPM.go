package PPM

import "fmt"

func (ppm *PPM) At(x, y int) Pixel {
	if y >= 0 && y < ppm.height && x >= 0 && x < ppm.width {
		fmt.Println("Choix du pixel: ", ppm.data[x][y])
		return ppm.data[x][y]
	}
	return Pixel{}
}

func (ppm *PPM) PrintData() {
	for x := 0; x < ppm.width; x++ {
		for y := 0; y < ppm.height; y++ {
			print(ppm.data[x][y].R, " ")
			print(ppm.data[x][y].G, " ")
			print(ppm.data[x][y].B, " ")
		}
		println()
	}
}

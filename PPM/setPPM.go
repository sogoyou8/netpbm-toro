package PPM

import "fmt"

func (ppm *PPM) Set(x, y int, value Pixel) {
	if y >= 0 && y < ppm.height && x >= 0 && x < ppm.width {
		fmt.Println("Valeur definit du pixel: ", value)
		//definitla valeur du pixel specifié
		ppm.data[x][y] = value
		return
	}
}

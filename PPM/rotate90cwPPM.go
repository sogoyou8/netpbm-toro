package PPM

func (ppm *PPM) Rotate90CW() {
	//inversee les lignes et les colones  pour 90wc
	ppm.width, ppm.height = ppm.height, ppm.height
	//criée un nouveau data
	newdata := make([][]Pixel, ppm.width)
	for i := range newdata {
		newdata[i] = make([]Pixel, ppm.height)
	}
	//remplir
	println("avant : ")
	for y := 0; y < ppm.height; y++ {
		for x := 0; x < ppm.width; x++ {
			print(ppm.data[x][y].R, " ")
			print(ppm.data[x][y].G, " ")
			print(ppm.data[x][y].B, " ")
		}
		println()
	}
	println("après")
	for y := 0; y < ppm.height; y++ {
		for x := 0; x < ppm.width; x++ {
			newdata[x][ppm.width-1-y] = ppm.data[y][x]
			// print du caractere + espace
			print(newdata[x][ppm.width-1-y].R, " ")
			print(newdata[x][ppm.width-1-y].G, " ")
			print(newdata[x][ppm.width-1-y].B, " ")

		}
		// retour a la ligne
		println()
	}
	ppm.data = newdata
}
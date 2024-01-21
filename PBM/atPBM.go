package PBM

import "fmt"

func (pbm *PBM) At(x, y int) bool {
	if y >= 0 && y < pbm.height && x >= 0 && x < pbm.width {
		fmt.Println("Choix du pixel: ", pbm.data[x][y])
		return pbm.data[x][y]
	}
	return false
}
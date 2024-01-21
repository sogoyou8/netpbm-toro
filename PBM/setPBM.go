package PBM

import "fmt"

func (pbm *PBM) Set(x, y int, value bool) {
	if y >= 0 && y < pbm.height && x >= 0 && x < pbm.width {
		fmt.Println("Valeur definit du pixel: ", value)
		//definitla valeur du pixel specifié
		pbm.data[y][x] = value
		return
	}

}

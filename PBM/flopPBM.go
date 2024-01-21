package PBM

import "fmt"

func (pbm *PBM) Flop() {
    for x := pbm.width - 1; x >= 0; x-- {
        for y := 0; y < pbm.height; y++ {
            if pbm.magicNumber == "P1" {
				if !pbm.data[x][y] {
					fmt.Print("0 ")
				} else {
					fmt.Print("1 ")
				}
			} else {
				if pbm.data[x][y] {
					fmt.Print("1")
				} else {
					fmt.Print("0")
				}
			}
        }
        fmt.Println()
    }
}
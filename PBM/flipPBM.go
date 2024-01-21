package PBM

import "fmt"

func (pbm *PBM) Flip() {
		for x := 0; x < pbm.width; x++ {
			for y := pbm.height - 1; y >= 0; y-- {
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
	
package PBM

import "fmt"

func (pbm *PBM) Invert() {
    for x := 0; x < pbm.width; x++ {
		for y := 0; y < pbm.height; y++ {
            if pbm.magicNumber == "P1" {
                if !pbm.data[x][y] {
                    fmt.Print("1 ")
                } else {
                    fmt.Print("0 ")
                }
            } else {
                if pbm.data[x][y] {
                    fmt.Print("0")
                } else {
                    fmt.Print("1")
                }
            }
        }
        fmt.Println()
    }
}

//inverse aussi a la verticale
func (pbm *PBM) InvertB() {
	for y := 0; y < pbm.height; y++ {
		for x := 0; x < pbm.width; x++ {
        switch pbm.magicNumber {
            case "P1":
                if pbm.data[x][y] {
                    fmt.Print("0 ")
                } else {
                    fmt.Print("1 ")
                }
            case "P4":
                if !pbm.data[x][y] {
                    fmt.Print("1")
                } else {
                    fmt.Print("0")
                }
        	}
    	}
    	fmt.Println()
	}
}
package PGM

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type PGM struct {
	data          [][]uint8
	width, height int
	magicNumber   string
	max           int
}

func parsePixel(value string) (bool, error) {
	switch value {
	case "0":
		return false, nil
	case "1":
		return true, nil
	default:
		return false, fmt.Errorf("unexpected pixel value: %s", value)
	}
}

// ReadPGM reads a PGM image from a file and returns a structure that represents the image
func ReadPGM(filename string) (*PGM, error) {
	// Open the PGM file
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %s", err)
	}
	defer file.Close()

	// Create a scanner for reading the file
	scanner := bufio.NewScanner(file)

	// Read the magic number
	var magicNumber string
	if !scanner.Scan() {
		return nil, fmt.Errorf("error reading magic number")
	}
	magicNumber = scanner.Text()

	// Read the width and height
	if !scanner.Scan() {
		return nil, fmt.Errorf("error reading width and height")
	}
	txt := scanner.Text()
	var width, height int
	fmt.Sscanf(txt, "%d %d", &width, &height)

	// Read the maximum value
	if !scanner.Scan() {
		return nil, fmt.Errorf("error reading max value")
	}
	max, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return nil, fmt.Errorf("error parsing max value: %s", err)
	}

	// Create the PGM struct
	pgm := &PGM{
		magicNumber: magicNumber,
		width:       width,
		height:      height,
		max:         max,
	}

	// Read the image data
	switch magicNumber {
	case "P2":
		// Read the image data as ASCII characters
		pgm.data = make([][]uint8, height)
		for y := 0; y < height; y++ {
			//println("y : ", y)
			if !scanner.Scan() {
				return nil, fmt.Errorf("error reading row %d", y)
			}
			row := scanner.Text()
			println(row)
			pgm.data[y] = make([]uint8, width)
			// variable ou on stockera la valeur en attente de conversion
			var value string
			// variable contenant la position du tableau à laquelle le nombre sera assigné
			var idx int
			for x, char := range row {
				// on vérifie si le caractere est numérique (entre 48 et 57 dans la table ascii)
				if char >= 48 && char <= 57 {
					// si c'est le cas alors on le stock dans la variable, cela nous permet de gérer les nombres a plusieurs chiffres
					value += string(char)
				} else {
					// quitte la boucle si le fichier est entierement parcouru
					if value == "" {
						break
					}
					// conversion en int
					valueInt, err := strconv.Atoi(value)
					if err != nil {
						return nil, fmt.Errorf("error converting pixel value at (%d, %d): %s", x, y, err)
					}
					// stockage dans le tableau
					pgm.data[y][idx] = uint8(valueInt)

					// reset de la variable value pour opérer sur un autre nombre
					value = ""

					// incrémentation d'idx afin pour affecter la prochaine valeur à la case suivante du tableau
					idx++
				}
				//println(" x : ", x, " char : ", char, " value : ", value)

			}
		}
	case "P5":
		// Read the image data as binary data
		binaryData, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, fmt.Errorf("error reading binary data: %s", err)
		}

		// Validate the image size
		expectedBytes := (width * height) * 3
		if len(binaryData) != expectedBytes {
			return nil, fmt.Errorf("error reading sizes, expected %d bytes, got %d bytes", expectedBytes, len(binaryData))
		}

		// Convert the binary data to uint8 values
		pgm.data = make([][]uint8, height)
		for y := 0; y < height; y++ {
			row := make([]uint8, width)
			for x := 0; x < width; x++ {
				byteIndex := y*width + x
				row[x] = binaryData[byteIndex*3]
			}
			pgm.data[y] = row
		}
	default:
	}
	return pgm, nil
}

package PPM

import (
	"bufio"
	"fmt"

	// "io/ioutil"
	"os"
	"strconv"
)

type PPM struct {
	data          [][]Pixel
	width, height int
	magicNumber   string
	max           uint
}

type Pixel struct {
	R, G, B uint8
}

// Readppm reads a ppm image from a file and returns a structure that represents the image
func ReadPPM(filename string) (*PPM, error) {
	// Open the ppm file
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
	maxUint := uint(max)

	// Create the ppm struct
	ppm := &PPM{
		magicNumber: magicNumber,
		width:       width,
		height:      height,
		max:         maxUint,
	}

	// Read the image data
	switch magicNumber {
	case "P3":
		// Read the image data as ASCII characters
		ppm.data = make([][]Pixel, height)
		for y := 0; y < height; y++ {
			if !scanner.Scan() {
				return nil, fmt.Errorf("error reading row %d", y)
			}
			row := scanner.Text()
			println(row)
			ppm.data[y] = make([]Pixel, width)
			// variable ou on stockera la valeur en attente de conversion
			var value string
			// variable contenant la position du tableau à laquelle le nombre sera assigné
			var idx int
			// variable qui contient l'info suivante : combien des 3 paramètres requis (R,G,B) ont été reçu
			var count int
			var red uint8
			var green uint8
			var blue uint8

			for x, char := range row {
				println("char : ", string(char))
				// on vérifie si le caractere est numérique (entre 48 et 57 dans la table ascii)
				if char >= 48 && char <= 57 {
					// si c'est le cas alors on le stock dans la variable, cela nous permet de gérer les nombres a plusieurs chiffres
					value += string(char)
				} else if char == 32 {
					// quitte la boucle si le fichier est entierement parcouru
					if value == "" {
						break
					}
					// conversion en int
					valueInt, err := strconv.Atoi(value)
					if err != nil {
						return nil, fmt.Errorf("error converting pixel value at (%d, %d): %s", x, y, err)
					}
					switch count {
					case 0:
						red = uint8(valueInt)
						value = ""
						count++
					case 1:
						green = uint8(valueInt)
						value = ""
						count++
					case 2:
						blue = uint8(valueInt)
						ppm.data[y][idx] = Pixel{R: red, G: green, B: blue}
						println("red : ", red, " green : ", green, " blue : ", blue, " IDX : ", idx)
						value = ""
						idx++
						count = 0
					}

				}
				//println(" x : ", x, " char : ", char, " value : ", value)

			}
		}
	// case "P6":
	// 	// Read the image data as binary data
	// 	binaryData, err := ioutil.ReadAll(file)
	// 	if err != nil {
	// 		return nil, fmt.Errorf("error reading binary data: %s", err)
	// 	}

	// 	// Validate the image size
	// 	expectedBytes := (width * height) * 3
	// 	if len(binaryData) != expectedBytes {
	// 		return nil, fmt.Errorf("error reading sizes, expected %d bytes, got %d bytes", expectedBytes, len(binaryData))
	// 	}

	// 	// Convert the binary data to uint8 values
	// 	ppm.data = make([][]Pixel, height)
	// 	for y := 0; y < height; y++ {
	// 		row := make([]Pixel, width)
	// 		// variable contenant la position du tableau à laquelle le nombre sera assigné
	// 		var idx int
	// 		// variable qui contient l'info suivante : combien des 3 paramètres requis (R,G,B) ont été reçu
	// 		var count int
	// 		var red uint8
	// 		var green uint8
	// 		var blue uint8
	// 		for x := 0; x < width; x++ {
	// 			byteIndex := y*width + x
	// 			row[x] = binaryData[byteIndex*3]
	// 		}
	// 		ppm.data[y] = row
	// 	}
	default:
	}
	return ppm, nil
}

package PGM

type PBM struct {
	data        [][]bool
	width       int
	height      int
	magicNumber string
}

func (pgm *PGM) ToPBM() *PBM {
	// Crée une matrice pour stocker les données de pgm (uint) en tant que pbm (bool)
	data := make([][]bool, pgm.height)
	for y := range data {
		data[y] = make([]bool, pgm.width)
	}
	height := pgm.height
	width := pgm.width
	magicNumber := "P1"

	// Détermine si un pixel sera noir ou blanc
	delta := pgm.max / 2

	// Remplit le tableau de données booléennes en fonction de la couleur du pixel (noir ou blanc)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if int(pgm.data[y][x]) <= delta {
				data[y][x] = false
			} else {
				data[y][x] = true
			}
		}
	}

	return &PBM{data, width, height, magicNumber}
}

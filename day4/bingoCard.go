package day4

type bingoCard struct {
	values       map[int]*valueMetadata
	rowsDrawn    [5]int
	columnsDrawn [5]int
	isBingo      bool
}

type valueMetadata struct {
	posRow    int
	posColumn int
	drawn     bool
}

func initializeCard(cardValues [5][5]int) *bingoCard {
	bc := &bingoCard{
		values:       map[int]*valueMetadata{},
		rowsDrawn:    [5]int{},
		columnsDrawn: [5]int{},
	}

	for row, rowValues := range cardValues {
		for column, value := range rowValues {
			bc.values[value] = &valueMetadata{
				posRow:    row,
				posColumn: column,
			}
		}
	}

	return bc
}

func (bc *bingoCard) draw(val int) {
	value, ok := bc.values[val]
	if !ok {
		return
	}

	value.drawn = true

	bc.rowsDrawn[value.posRow]++
	if bc.rowsDrawn[value.posRow] == 5 {
		bc.isBingo = true
	}

	bc.columnsDrawn[value.posColumn]++
	if bc.columnsDrawn[value.posColumn] == 5 {
		bc.isBingo = true
	}
}

func (bc bingoCard) undrawnValue() int {
	totalValue := 0

	for val, meta := range bc.values {
		if meta.drawn == false {
			totalValue += val
		}
	}

	return totalValue
}

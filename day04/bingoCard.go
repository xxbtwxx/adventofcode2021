package day04

const (
	rows    = 5
	columns = 5
)

type bingoCard struct {
	values       map[int]*valueMetadata
	rowsDrawn    [rows]int
	columnsDrawn [columns]int
	isBingo      bool
}

type valueMetadata struct {
	posRow    int
	posColumn int
	drawn     bool
}

func initializeCard(cardValues [rows][columns]int) *bingoCard {
	bc := &bingoCard{
		values:       map[int]*valueMetadata{},
		rowsDrawn:    [rows]int{},
		columnsDrawn: [columns]int{},
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

func (bc *bingoCard) mark(val int) {
	value, ok := bc.values[val]
	if !ok {
		return
	}

	value.drawn = true

	bc.rowsDrawn[value.posRow]++
	if bc.rowsDrawn[value.posRow] == rows {
		bc.isBingo = true
	}

	bc.columnsDrawn[value.posColumn]++
	if bc.columnsDrawn[value.posColumn] == columns {
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

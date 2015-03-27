package dot

import "bytes"

// bitVals contains the bitwise values used to generate the Unicode offset.
var bitVals = [4][2]byte{
	{0x01, 0x08},
	{0x02, 0x10},
	{0x04, 0x20},
	{0x40, 0x80},
}

func Render(pix [][]bool) string {
	if pix == nil {
		return ""
	}
	buf := bytes.Buffer{}
	first := true
	rowOffset := 0
	curRow := [][4][2]bool{}
	for row := range pix {
		// Grow the current row if it's not long enough.
		if l := (len(pix[row]) + 1) / 2; l > len(curRow) {
			newRow := make([][4][2]bool, l)
			copy(newRow, curRow)
			curRow = newRow
		}
		// Set the pixels in the current row.
		for i, p := range pix[row] {
			curRow[i/2][rowOffset][i%2] = p
		}
		// Prepare for next row, writing and creating a new one if needed.
		rowOffset++
		if rowOffset == 4 {
			if !first {
				buf.WriteByte('\n')
			}
			buf.WriteString(RuneRow(curRow))

			first = false
			rowOffset = 0
			curRow = [][4][2]bool{}
		}
	}
	if rowOffset > 0 {
		if !first {
			buf.WriteByte('\n')
		}
		buf.WriteString(RuneRow(curRow))
	}
	return buf.String()
}

func Rune(pix [4][2]bool) rune {
	offset := byte(0)
	for row := range pix {
		for col, on := range pix[row] {
			if on {
				offset += bitVals[row][col]
			}
		}
	}
	return '⠀' + rune(offset)
}

func RuneRow(row [][4][2]bool) string {
	if row == nil {
		return ""
	}
	runes := make([]rune, len(row))
	for i, pix := range row {
		runes[i] = Rune(pix)
	}
	return string(runes)
}

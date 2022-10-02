package main

import "strings"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	step := numRows*2 - 2
	b := strings.Builder{}
	for i := 0; i < numRows; i++ {
		between := i > 0 && i < numRows-1
		for j := i; j < len(s); j += step {
			b.WriteByte(s[j])
			if between && j+step-2*i < len(s) {
				b.WriteByte(s[j+step-2*i])
			}
		}
	}
	return b.String()
}

func convert2(s string, numRows int) string {
	sheet := make([][]byte, numRows)
	index := 0
	back := false
	for i := range s {
		if back && index <= 1 {
			back = false
			index = 0
		}
		if !back {
			sheet[index] = append(sheet[index], s[i])
			index++
			if index == numRows {
				back = true
				index--
			}
			continue
		}
		if index > 0 {
			index--
			sheet[index] = append(sheet[index], s[i])
		}
	}
	b := strings.Builder{}
	for i := range sheet {
		b.Write(sheet[i])
	}
	return b.String()
}

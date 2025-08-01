package main

import (
	"bytes"
	"strconv"
	"strings"
)

func comma(f float64) string {
	s := strconv.FormatFloat(f, 'f', -1, 64)
	parts := strings.Split(s, ".")
	integerPart := parts[0]
	decimalPart := ""
	if len(parts) > 1 {
		decimalPart = "." + parts[1]
	}
	n := len(integerPart)
	if n <= 3 {
		return integerPart + decimalPart
	}
	var buf bytes.Buffer
	firstGroupSize := n % 3
	if firstGroupSize == 0 {
		firstGroupSize = 3
	}
	buf.WriteString(integerPart[:firstGroupSize])
	for i := firstGroupSize; i < n; i += 3 {
		buf.WriteByte(',')
		buf.WriteString(integerPart[i : i+3])
	}
	return buf.String() + decimalPart
}

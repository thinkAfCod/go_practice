package main

import (
	"strconv"
)

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}
	result := ""
	if (numerator < 0 && denominator > 0) || (numerator > 0 && denominator < 0) {
		result += "-"
	}
	nt, dt := abs(numerator), abs(denominator)
	remained := nt % dt
	result += strconv.FormatInt(int64(nt/dt), 10)
	if remained == 0 {
		return result
	}
	result += "."
	i := 0
	divResMap := make(map[int]int)
	for remained != 0 {
		inner := remained * 10 / dt

		if divResMap[remained] != 0 {
			result += "("
			result += strconv.Itoa(divResMap[remained])
			result += ")"
			break
		}
		i++
		divResMap[remained] = i
		result += strconv.Itoa(inner)
		remained = (remained * 10) % dt
	}
	return result
}

func abs(n int) int {
	y := n >> 31
	return (n ^ y) - y
}

func main() {
	println(fractionToDecimal(2346245, 1346245))
}

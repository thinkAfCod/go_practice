package main

import (
	"fmt"
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	ver1Arr := strings.Split(version1, ".")
	ver2Arr := strings.Split(version2, ".")
	ver1Lenght := len(ver1Arr)
	ver2Lenght := len(ver2Arr)
	for i := 0; i < ver1Lenght || i < ver2Lenght; i++ {
		digit1 := 0
		digit2 := 0
		if i < ver1Lenght {
			digit1, _ = strconv.Atoi(ver1Arr[i])
		}
		if i < ver2Lenght {
			digit2, _ = strconv.Atoi(ver2Arr[i])
		}
		if digit1 > digit2 {
			return 1
		} else if digit1 < digit2 {
			return -1
		}
	}
	return 0
}

func main() {
	version := compareVersion("1.1.1", "1.0001")
	fmt.Println(version)
}

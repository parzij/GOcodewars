package kata

import (
	"fmt"
	"strings"
	"strconv"
)

func HighAndLow(in string) string {
	parts := strings.Split(in, " ")

	maxEl, _ := strconv.Atoi(parts[0])
	minEl := maxEl

	for _, s := range parts {
		num, _ := strconv.Atoi(s)
		if num > maxEl {
			maxEl = num
		}
		if num < minEl {
			minEl = num
		}
	}

	return fmt.Sprintf("%d %d", maxEl, minEl)
}

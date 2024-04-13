package calculations

import (
	"log"
	"strconv"
	"strings"
)

func ConvertToFloat(value []string, index int) float64 {
	value[index] = strings.TrimSpace(value[index])
	digit, err := strconv.ParseFloat(value[index], 64)
	if err != nil {
		log.Fatal(err)
	}
	return digit

}

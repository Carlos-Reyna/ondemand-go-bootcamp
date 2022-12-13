package utils

import (
	"log"
	"strconv"
)

func ToInt(s string) int {
	val, err := strconv.Atoi(s)

	if err != nil {
		log.Fatal(err.Error())
	}

	return val
}

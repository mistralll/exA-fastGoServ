package main

import (
	"fmt"
	"log"

	"github.com/mistralll/expAcsv/mergeCsvUtil"
)

func main() {
	line, err := mergeCsvUtil.GetGeotagLine("2054419499")
	if err != nil {
		log.Fatal(err)
	}

	for _, row := range line {
		fmt.Print(row + " ")
	}
}

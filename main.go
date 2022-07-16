package main

import (
	"log"

	"github.com/mistralll/expAcsv/csvComp"
)

func main() {
	// serv.TagData = make([]serv.Tag, 6736297)
	// err := serv.ReadData("csv/comp/merged.csv")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// serv.ServRun()

	err := csvComp.CntTagToCsv("csv/comp/tagComped.csv", "csv/cnt/cnt.csv")
	if err != nil {
		log.Fatal(err)
	}

}

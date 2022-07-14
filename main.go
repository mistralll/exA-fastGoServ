package main

import (
	"log"

	"github.com/mistralll/expAcsv/serv"
)

func main() {
	serv.TagData = make([]serv.Tag, 6736297)
	err := serv.ReadData("csv/comp/merged.csv")
	if err != nil {
		log.Fatal(err)
	}
	// serv.ServRun()

	serv.PrintTagData()

}

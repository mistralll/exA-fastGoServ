package main

import (
	"fmt"
	"log"

	"github.com/mistralll/expAcsv/serv"
)

func main() {
	serv.TagData = make([]serv.Tag, 860621)
	err := serv.ReadData("csv/cmp/mrg.csv", "csv/cmp/cnt.csv")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("server is running...")
	serv.ServRun()
}

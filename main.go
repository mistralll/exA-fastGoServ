package main

import (
	"github.com/mistralll/expAcsv/csvUtil"
)

func main() {
	csvUtil.AddInf("csv/sorted/tag.csv", "csv/sorted/geotag.csv", "csv/marged/data.csv")
}

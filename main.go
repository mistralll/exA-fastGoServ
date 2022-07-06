package main

import (
	"github.com/mistralll/expAcsv/csvUtil"
)

func main() {
	csvUtil.Merge("csv/sorted/tag.csv", "csv/sorted/geotag.csv", "csv/merged/data.csv")
}

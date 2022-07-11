package main

import (
	"github.com/mistralll/expAcsv/csvUtil"
)

func main() {
	csvUtil.DelEmptyTag("csv/merged/data.csv", "csv/comp/data.csv")
}

package main

import (
	"github.com/mistralll/expAcsv/csvUtil"
)

func main() {
	csvUtil.DelEmptyTag("csv/comp/data.csv", "csv/comp/sorted.csv")
}

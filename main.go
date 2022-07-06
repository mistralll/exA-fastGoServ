package main

import "github.com/mistralll/expAcsv/tagUtil"

func main() {
	tagUtil.SortCsv("csv/rowdata/tag.csv", "csv/sorted/tag.csv")
}

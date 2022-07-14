package main

import "github.com/mistralll/expAcsv/csvComp"

func main() {
	// serv.ReadData("csv/comp/cmptag01.csv", "csv/comp/cmpimg01.csv")
	// serv.ServRun()

	csvComp.CompSortedTagData("csv/sorted/tagSortedByTagDate.csv", "csv/comp/tagComped.csv")
}

package main

import (
	"github.com/mistralll/expAcsv/serv"
)

func main() {
	serv.ReadData("csv/comp/cmptag01.csv", "csv/comp/cmpimg01.csv")
	serv.ServRun()
}

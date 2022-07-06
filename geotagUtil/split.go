package geotagUtil

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

func SplitCsv(index int) {
	ans := [][]string{}

	f, err := os.Open("geotag/geotag.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	r := csv.NewReader(f)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		recstr := record[0]

		recint, _ := strconv.Atoi(recstr)

		iii := recint % 100

		if iii == index {
			ans = append(ans, record)
		}

	}

	csvout, err := os.Create("data/geotag-" + strconv.Itoa(index) + ".csv")
	if err != nil {
		log.Fatal(err)
	}
	defer csvout.Close()

	w := csv.NewWriter(csvout)

	for _, rec := range ans {
		if err := w.Write(rec); err != nil {
			log.Fatal(err)
		}
	}

	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}

}

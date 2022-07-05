package csvUtil

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"sort"
)

func SortCsv(inFileName string, outFileName string) {
	infile, err := os.Open(inFileName)
	if err != nil {
		log.Fatal(err)
	}

	defer infile.Close()

	r := csv.NewReader(infile)

	vec := [][]string{}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		vec = append(vec, record)
	}

	sort.SliceStable(vec, func(i, j int) bool {
		if vec[i][0] != vec[j][0] {
			return vec[i][0] < vec[j][0]
		} else {
			return vec[i][1] < vec[j][1]
		}
	})

	outfile, err := os.Create(outFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer outfile.Close()

	w := csv.NewWriter(outfile)

	for _, record := range vec {
		if err := w.Write(record); err != nil {
			log.Fatal(err)
		}
	}

	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}

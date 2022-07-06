package geotagUtil

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func CheckAll() {
	for i := 0; i < 100; i++ {
		filename := "sorted/geotag-" + strconv.Itoa(i) + ".csv"
		err := CheckCsvDoubling(filename)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func CheckCsvDoubling(infilieName string) error {
	infile, err := os.Open(infilieName)
	if err != nil {
		return err
	}

	defer infile.Close()

	r := csv.NewReader(infile)

	vec := [][]string{}

	cnt := 0
	i := 0

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		vec = append(vec, record)

		if i != 0 {
			if vec[i-1][0] == vec[i][0] {
				cnt++
			} else {
				if(cnt != 0) {
					fmt.Print(cnt)
					cnt = 0
				}
			}
		}

		i++
	}

	return nil
}

package mergeCsvUtil

import (
	"encoding/csv"
	"io"
	"math"
	"os"
)

func GetGeotagLine(id string) ([]string, error) {
	// set file name
	fileno := id[8:10]
	filename := "csv/sortedGeotag/geotag-" + fileno + ".csv"

	// open file
	infile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer infile.Close()

	r := csv.NewReader(infile)

	// read to vec
	vec := [][]string{}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		vec = append(vec, record)
	}

	// binary search
	ok := len(vec)
	ng := -1
	mid := (ok + ng) / 2
	for math.Abs(float64(ok-ng)) > 1 {
		mid = (ok + ng) / 2
		if vec[mid][0] < id {
			ng = mid
		} else {
			ok = mid
		}
	}

	return vec[ok], nil

}

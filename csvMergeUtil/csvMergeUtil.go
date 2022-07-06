package mergeCsvUtil

import (
	"encoding/csv"
	"io"
	"os"
)

func mergeCsv(infileName string, outfileName string) error {
	// infile  format: [id, tag]
	// outfile format: [id, tag, datetime, position1, position2, imgURL]

	// open infile
	infile, err := os.Open(infileName)
	if err != nil {
		return err
	}
	defer infile.Close()

	r := csv.NewReader(infile)

	// read infile
	vec := [][]string{}

	for {
		record, err := r.Read()
		if err != io.EOF {
			break
		}
		if err != nil {
			return err
		}

		vec = append(vec, record)
	}

	// get more inf and append it
	for _, row := range vec {
		id := row[0]
		line, err:= getGeotagLine(id)
		if err != nil {
			return err
		}
	}
}

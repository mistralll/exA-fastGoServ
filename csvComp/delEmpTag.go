package csvComp

import (
	"encoding/csv"
	"io"
	"os"
)

func DelEmpTag(infileName string, outfileName string) error {
	/*
		Input  CSV file: [id, tag]
		Output CSV file: [id, tag]

		* if id is empty, its line is not include to output csv file.
	*/

	// read file
	infile, err := os.Open(infileName)
	if err != nil {
		return err
	}
	defer infile.Close()

	// if not "", then append row
	ans := [][]string{}
	r := csv.NewReader(infile)

	for {
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		if rec[1] != "" {
			ans = append(ans, rec)
		}
	}

	// write
	outfile, err := os.Create(outfileName)
	if err != nil {
		return err
	}
	defer outfile.Close()

	w := csv.NewWriter(outfile)

	for _, rec := range ans {
		if err := w.Write(rec); err != nil {
			return err
		}
	}

	w.Flush()

	return nil
}

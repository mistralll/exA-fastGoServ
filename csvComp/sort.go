package csvComp

import (
	"encoding/csv"
	"io"
	"os"
	"sort"
)

func SortTagById(infileName string, outfileName string) error {
	// read file
	infile, err := os.Open(infileName)
	if err != nil {
		return err
	}
	defer infile.Close()

	// infile to vec: tag
	r := csv.NewReader(infile)
	tag := [][]string{}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		tag = append(tag, record)
	}

	// sort
	sort.SliceStable(tag, func(i, j int) bool {
		return tag[i][0] < tag[j][0] 
	})

	// create outfile
	outfile, err := os.Create(outfileName)
	if err != nil {
		return err
	}
	defer outfile.Close()

	// write to outfile
	w := csv.NewWriter(outfile)

	for _, record := range tag {
		if err := w.Write(record); err != nil {
			return err
		}
	}

	w.Flush()

	if err := w.Error(); err != nil {
		return err
	}

	return nil
}

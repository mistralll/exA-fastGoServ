package csvComp

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
)

func CntTagToCsv(tagfileName string, outfileName string) error {
	tagfile, err := os.Open(tagfileName)
	if err != nil {
		return err
	}
	defer tagfile.Close()

	// infile to vec: tag
	r := csv.NewReader(tagfile)
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

	// cntTable
	cntTable := [][]string{}
	cnt := 0
	for _, row := range tag {
		cnt++
		if row[0] != "" {
			cntTable = append(cntTable, []string{strconv.Itoa(cnt)})
			cnt = 0
		}
	}

	// create outfile
	outfile, err := os.Create(outfileName)
	if err != nil {
		return err
	}
	defer outfile.Close()

	// write to outfile
	w := csv.NewWriter(outfile)

	for _, record := range cntTable {
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

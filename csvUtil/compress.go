package csvUtil

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

func DelEmptyTag(infileName string, outfileName string) error {
	// read infile
	infile, err := os.Open(infileName)
	if err != nil {
		return err
	}
	defer infile.Close()

	r := csv.NewReader(infile)

	// if tag is not empty, then append to vec: data
	data := [][]string{}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		if record[1] != "" {
			data = append(data, record)
		}

		if len(data)%1000000 == 0 && len(data) > 0 {
			fmt.Println(len(data))
		}
	}

	fmt.Println("check")

	// wirte vec: data to csv file
	outfile, err := os.Create(outfileName)
	if err != nil {
		return err
	}
	defer outfile.Close()

	w := csv.NewWriter(outfile)

	for i, record := range data {
		if err := w.Write(record); err != nil {
			return err
		}

		if i % 1000000 == 0 {
			fmt.Println(strconv.Itoa(i) + " / " + strconv.Itoa(len(data)))
		}
	}

	w.Flush()

	if err := w.Error(); err != nil {
		return err
	}

	return nil
}

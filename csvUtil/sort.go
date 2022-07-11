package csvUtil

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

func SortTagAndDate(infileName string, outfileName string) error {
	// read file
	infile, err := os.Open(infileName)
	if err != nil {
		return err
	}
	defer infile.Close()

	r := csv.NewReader(infile)

	// infile to vec: data
	data := [][]string{}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		data = append(data, record)

		if len(data)%100000 == 0 && len(data)>0 {
			fmt.Print("csvUtil.SortTagDate: " + strconv.Itoa(len(data)));
		}
	}

	// sort
	sort.SliceStable(data, func(i, j int) bool {
		if data[i][1] != data[j][1] {
			return data[i][1] < data[j][1]
		} else {
			return data[i][2] < data[j][2]
		}
	})

	// create new file
	outfile, err := os.Create(outfileName)
	if err != nil {
		return err
	}
	defer outfile.Close()

	w := csv.NewWriter(outfile)

	// write to file
	for i, record := range data {
		if err := w.Write(record); err != nil {
			return err
		}

		if i % 10000 == 0 {
			fmt.Println("csvUtil.SortTagDate: " + strconv.Itoa(i) + " / " + strconv.Itoa(len(data)))
		}
	}

	return nil
}

package csvComp

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"sort"
	"time"
)

func SortByFirstColumm(infileName string, outfileName string) error {
	/*
		This func sort order by first columm
	*/

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

func SortTagWithDateByTagDate(infileNmae string, outfileName string) error {
	/*
		csv format: [id, tag, date]
		sort by tag and date
	*/

	fmt.Println("now loading...")
	// read tag file
	tagfile, err := os.Open(infileNmae)
	if err != nil {
		return err
	}
	defer tagfile.Close()

	// tagfile to vec: tag
	r := csv.NewReader(tagfile)
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
	}

	fmt.Println("now sorting...")

	// sort
	sort.SliceStable(data, func(i, j int) bool {
		if data[i][1] != data[j][1] {
			return data[i][1] < data[j][1]
		} else {
			iTime, _ := time.Parse("2006-01-02 15:04:05", data[i][2])
			jTime, _ := time.Parse("2006-01-02 15:04:05", data[j][2])

			return iTime.After(jTime)
		}
	})

	fmt.Println("now writing...")

	// write
	outfile, err := os.Create(outfileName)
	if err != nil {
		return err
	}
	defer outfile.Close()

	w := csv.NewWriter(outfile)

	// write to file
	for _, record := range data {
		if err := w.Write(record); err != nil {
			return err
		}
	}

	w.Flush()

	return nil
}

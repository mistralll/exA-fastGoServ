package tagUtil

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

func SortAllCsv() {
	table := "abcdefghijklmnopqrstuvwxyz"
	for i := 0; i < 26; i++ {
		infile := "csv/splitedTag/tag-" + table[i:i+1] + ".csv"
		outfile := "csv/sortedTag/tag-" + table[i:i+1] + ".csv"
		SortCsv(infile, outfile)
		fmt.Println(table[i:i+1] + " is complited!")
	}

	infile := "csv/splitedTag/tag-#.csv"
	outfile := "csv/sortedTag/tag-#.csv"
	SortCsv(infile, outfile)
	fmt.Println("# is complited!")

}

func SortCsv(infileName string, outfileName string) error {
	infile, err := os.Open(infileName)
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
			return err
		}

		vec = append(vec, record)
	}

	sort.SliceStable(vec, func(i, j int) bool { return vec[i][1] < vec[j][1] })

	outfile, err := os.Create(outfileName)
	if err != nil {
		return err
	}
	defer outfile.Close()

	w := csv.NewWriter(outfile)

	for _, record := range vec {
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

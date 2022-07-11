package tagUtil

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

/*
	legacy code!
	このコードは方針転換前のもはや使い道のないコードですが設計の参考に残してあります。
*/

func SplitAll() {
	SplitCsv("csv/rowdata/tag.csv", "csv/splitedTag/tag-#.csv", 27)
}

func SplitCsv(infileName string, outfileName string, index int) error {

	// prepare file
	infile, err := os.Open(infileName)
	if err != nil {
		return err
	}
	defer infile.Close()

	r := csv.NewReader(infile)

	// make ans vector
	ans := [][]string{}

	// make key as a slice
	table := "abcdefghijklmnopqrstuvwxyz#"

	for {
		// read a line
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		tag := record[1]

		if tag == "" {
			continue
		}

		if index == 27 {
			// tag[0:1] != a ~ z であることを確認したら append する
			flag := true
			for i := 0; i < 26; i++ {
				if tag[0:1] == table[i:i+1] {
					flag = false
					break
				}
			}
			if flag {
				ans = append(ans, record)
			}

		} else {
			if tag[0:1] == table[index:index+1] {
				ans = append(ans, record)
			}
		}
	}

	outfile, err := os.Create(outfileName)
	if err != nil {
		return err
	}
	defer outfile.Close()

	w := csv.NewWriter(outfile)

	for _, rec := range ans {
		if err := w.Write(rec); err != nil {
			log.Fatal(err)
		}
	}

	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}

	return nil
}

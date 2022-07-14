package csvComp

import (
	"encoding/csv"
	"io"
	"os"
)

func CompSortedTagData(infileName string, outfileName string) error {
	/*
		input csv : [id, tag, date]
		output csv: [tag, id]

		入力のcsvに関しては tag, date の順にソートされていること。

		タグが空白の行は出力しない。
		同じタグが100件以上あった場合、はじめの行から100個のみ出力する。
		同じタグが繰り返される場合、2行目以降はタグを""にする。
	*/

	// read file
	infile, err := os.Open(infileName)
	if err != nil {
		return err
	}
	defer infile.Close()

	// infile to vec: tag
	tag := [][]string{}
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
			tag = append(tag, rec)
		}
	}

	// make vec: ans
	ans := [][]string{}
	cnt := 0
	prev := ""
	for _, row := range tag {
		if row[1] == prev {
			if cnt < 100 {
				ans = append(ans, []string{"", row[0]})
			}
		} else {
			cnt = 0
			ans = append(ans, []string{row[1], row[0]})
		}
		prev = row[1]
		cnt++
	}

	// write filie
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

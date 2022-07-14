package csvComp

import (
	"encoding/csv"
	"io"
	"math"
	"os"
)

func CompSortedTagData (infileName string, outfileName string) error{
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

func DelEmpTag(infileName string, outfileName string) error{
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
		if err != nil{
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

func DelEmpTagImg(tagFileTitle string, imgFileTitle string, outImgFile string) error {
/*
	need to remake this function.
*/


	// read tag file
	tagFile, err := os.Open(tagFileTitle)
	if err != nil {
		return err
	}
	defer tagFile.Close()

	// tag.csv to vec: tag
	tagReader := csv.NewReader(tagFile)
	tag := [][]string{}

	for {
		record, err := tagReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		tag = append(tag, record)
	}

	// read img file
	imgFile, err := os.Open(imgFileTitle)
	if err != nil {
		return err
	}
	defer imgFile.Close()

	// img file to vec: img
	imgReader := csv.NewReader(imgFile)
	img := [][]string{}

	for {
		record, err := imgReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		img = append(img, record)
	}

	ans := [][]string{}
	// img を１つずつ見ていき、tagを探索。一致したらansにappend
	for _, row := range img {
		id := row[0]

		// binary serch
		ok := len(tag)
		ng := -1
		for math.Abs(float64(ok-ng)) > 1 {
			mid := (ok + ng) / 2
			if tag[mid][0] < id {
				ng = mid
			} else {
				ok = mid
			}
		}

		if len(tag) != ok {
			if tag[ok][1] != "" {
				ans = append(ans, row)
			}
		}
	}

	// write file
	outfile, err := os.Create(outImgFile)
	if err != nil {
		return err
	}
	defer outfile.Close()

	w := csv.NewWriter(outfile)

	for _, record := range ans {
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

package csvComp

import (
	"encoding/csv"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

func SplitURLSub(URL string) [3]string {
	var ans [3]string
	ans[0] = URL[11:12]

	// find 3rd and 4th '/' index
	index3 := -1
	index4 := -1
	cnt := 0
	for i, c := range URL {
		if c == '/' {
			cnt++
		}
		if cnt == 3 {
			index3 = i
			break
		}
	}
	cnt = 0
	for i, c := range URL {
		if c == '/' {
			cnt++
		}
		if cnt == 4 {
			index4 = i
			break
		}
	}
	ans[1] = URL[index3+1 : index4]

	// find _
	index5 := -1
	for i, c := range URL {
		if c == '_' {
			index5 = i
			break
		}
	}
	ans[2] = URL[index5+1 : len(URL)-4]
	return ans
}

func SplitURLOfImg(infileName string, outfileName string) error {
	// read file
	infile, err := os.Open(infileName)
	if err != nil {
		return err
	}
	defer infile.Close()

	r := csv.NewReader(infile)
	data := [][]string{}

	for {
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		// ここで整形
		new := SplitURLSub(rec[4])

		data = append(data, []string{rec[0], rec[1], rec[2], rec[3], new[0], new[1], new[2]})
	}

	outfile, err := os.Create(outfileName)
	if err != nil {
		return err
	}
	defer outfile.Close()

	w := csv.NewWriter(outfile)

	for _, rec := range data {
		if err := w.Write(rec); err != nil {
			return err
		}
	}

	w.Flush()

	if err := w.Error(); err != nil {
		return err
	}

	return nil
}

func DelNotUseRow(tagfileName string, imgfileName string, outfileName string) error {
	/*
		imgを一行ずつみていき、tagのなかにidが含まれていたらappendします。
		tagの探索に二分探索を使用しているので、tagファイルは第一要素のidをソートしておいてください。
		tagファイルにidを見つけたら即appendしているので、空白タグに関しては事前に削除してください。
	*/

	// read tag file
	fmt.Println("loading tagfile...")
	tagFile, err := os.Open(tagfileName)
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
	fmt.Println("loading img file...")
	imgFile, err := os.Open(imgfileName)
	if err != nil {
		return err
	}
	defer imgFile.Close()

	// img file to vec: img
	imgReader := csv.NewReader(imgFile)
	img := [][]string{}

	cnt := 0

	for {
		record, err := imgReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		// binary search
		ok := len(tag)
		ng := -1
		key := record[0]
		for math.Abs(float64(ok-ng)) > 1 {
			mid := (ok + ng) / 2
			if tag[mid][1] < key {
				ng = mid
			} else {
				ok = mid
			}
		}

		if 0 <= ok && ok < len(tag) {
			img = append(img, record)
		}

		// print progres
		if cnt%100000 == 0 {
			fmt.Println(strconv.Itoa(cnt))
		}

		cnt++
	}

	// write to csv
	outfile, err := os.Create(outfileName)
	if err != nil {
		return err
	}
	defer outfile.Close()

	w := csv.NewWriter(outfile)

	for _, rec := range img {
		if err := w.Write(rec); err != nil {
			return err
		}
	}

	w.Flush()

	if err := w.Error(); err != nil {
		return err
	}

	return nil
}

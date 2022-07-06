package csvUtil

import (
	"encoding/csv"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

func Merge(tagFileName string, imgFileName string, outfileName string) error {
	// read tag.csv
	tagFile, err := os.Open(tagFileName)
	if err != nil {
		return err
	}
	defer tagFile.Close()

	tagReader := csv.NewReader(tagFile)

	// tag.csv to vec
	tags := [][]string{}

	for {
		record, err := tagReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		tags = append(tags, record)
	}

	// read geotag.csv
	imgFile, err := os.Open(imgFileName)
	if err != nil {
		return err
	}
	defer imgFile.Close()

	imgReader := csv.NewReader(imgFile)

	// geotag.csv to vec
	imgs := [][]string{}

	for {
		record, err := imgReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		imgs = append(imgs, record)
	}

	// open new file
	outfile, err := os.Create(outfileName)
	if err != nil {
		return err
	}
	defer outfile.Close()

	outWriter := csv.NewWriter(outfile)

	// add inf to vec of tags
	for i, row := range tags {
		key := row[0]

		// binary search
		ok := len(imgs)
		ng := -1
		for math.Abs(float64(ok-ng)) > 1 {
			mid := (ok + ng) / 2
			if imgs[mid][0] < key {
				ng = mid
			} else {
				ok = mid
			}
		}

		ans := []string{tags[i][0], tags[i][1], imgs[ok][1], imgs[ok][2], imgs[ok][3], imgs[ok][4]}

		if err := outWriter.Write(ans); err != nil {
			return err
		}

		outWriter.Flush()

		if err := outWriter.Error(); err != nil {
			return err
		}

		if i%100 == 0 {
			fmt.Println(strconv.Itoa(i) + " / " + strconv.Itoa(len(tags)))
		}
	}

	return nil

}

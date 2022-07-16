package csvComp

import (
	"encoding/csv"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

func MergeCsvToCompedTag(tagfileName string, imgfileNmae string, outfileName string) error {
/*
	input tag file: [tag, id]
	input img file: [id, date, loc1, loc2, URL1, URL2, URL3]
	output merge file: [tag, id, date, loc1, loc2, URL1, URL2, URL3]
*/

	// read tag.csv
	tagFile, err := os.Open(tagfileName)
	if err != nil {
		return err
	}
	defer tagFile.Close()

	tagReader := csv.NewReader(tagFile)

	// tag.csv to vec
	tagdata := [][]string{}

	for {
		record, err := tagReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		tagdata = append(tagdata, record)
	}
	fmt.Println("tag is read!")
	// read geotag.csv
	imgFile, err := os.Open(imgfileNmae)
	if err != nil {
		return err
	}
	defer imgFile.Close()

	imgReader := csv.NewReader(imgFile)

	// geotag.csv to vec
	imgdata := [][]string{}

	for {
		record, err := imgReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		imgdata = append(imgdata, record)
	}

	// open new file
	outfile, err := os.Create(outfileName)
	if err != nil {
		return err
	}
	defer outfile.Close()

	outWriter := csv.NewWriter(outfile)

	// add inf to vec of tags
	for i, row := range tagdata {
		key := row[1]

		// binary search
		ok := len(imgdata) - 1
		ng := 0
		for math.Abs(float64(ok-ng)) > 1 {
			mid := (ok + ng) / 2
			if imgdata[mid][0] < key {
				ng = mid
			} else {
				ok = mid
			}
		}

		// make new line and write it.
		ans := []string{tagdata[i][0], tagdata[i][1], imgdata[ok][1], imgdata[ok][2], imgdata[ok][3], imgdata[ok][4], imgdata[ok][5], imgdata[ok][6]}

		if err := outWriter.Write(ans); err != nil {
			return err
		}

		outWriter.Flush()

		if err := outWriter.Error(); err != nil {
			return err
		}

		// print progress
		if i%100 == 0 {
			fmt.Println(strconv.Itoa(i) + " / " + strconv.Itoa(len(tagdata)))
		}
	}

	return nil

}

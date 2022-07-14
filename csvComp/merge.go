package csvComp

import (
	"encoding/csv"
	"fmt"
	"io"
	"math"
	"os"
)

/*
	this funcs are make for tag.csv
*/

func AddDateInfToTag(tagfileName string, imgfileName string, outfilieName string) error {

	fmt.Println("now loading...")
	// read tag file
	tagfile, err := os.Open(tagfileName)
	if err != nil {
		return err
	}
	defer tagfile.Close()

	// tagfile to vec: tag
	r := csv.NewReader(tagfile)
	tg := [][]string{}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		tg = append(tg, record)
	}

	fmt.Println("now loading...")
	// read img file
	imgfile, err := os.Open(imgfileName)
	if err != nil {
		return err
	}
	defer imgfile.Close()

	// imgfile to vec: img
	r = csv.NewReader(imgfile)
	im := [][]string{}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		im = append(im, record)
	}

	// add inf
	for cnt, row := range tg {
		// binary serch
		key := row[0]
		ok := len(im)
		ng := -1
		for math.Abs(float64(ok-ng)) > 1 {
			mid := (ok + ng) / 2
			if im[mid][0] < key {
				ng = mid
			} else {
				ok = mid
			}
		}

		// add inf
		tg[cnt] = append(tg[cnt], im[ok][1])

		if cnt%10000 == 0 {
			fmt.Println(cnt)
		}
	}

	fmt.Print("Check")

	// write to outfile
	outfile, err := os.Create(outfilieName)
	if err != nil {
		return err
	}
	w := csv.NewWriter(outfile)

	for _, record := range tg {
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

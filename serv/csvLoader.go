package serv

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"runtime"
	"strconv"
)

func ReadData(dataName string, cntName string) error {
	fmt.Println("csvLoader: now reading data...")

	// read cnt file
	cntfile, err := os.Open(cntName)
	if err != nil {
		return err
	}
	defer cntfile.Close()
	r := csv.NewReader(cntfile)

	cnt := [860622]int32{}
	i := 0
	for {
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		tmp, err := strconv.Atoi(rec[0])
		if err != nil {
			return err
		}

		cnt[i] = int32(tmp)

		i++
	}

	// read tag file
	infile, err := os.Open(dataName)
	if err != nil {
		return err
	}
	defer infile.Close()
	r = csv.NewReader(infile)

	tagcnt := -1
	imgcnt := 0

	for {
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		if rec[0] != "" {
			tagcnt++
			imgcnt = 0
			TagData[tagcnt].Name = rec[0]
			TagData[tagcnt].Imgs = make([]Image, cnt[tagcnt])

			if tagcnt%2000 == 0 {
				runtime.GC()
				fmt.Println(tagcnt)
			}
		}
		id, err := strconv.ParseInt(rec[1], 10, 64)
		if err != nil {
			return err
		}
		TagData[tagcnt].Imgs[imgcnt].Id = id
		TagData[tagcnt].Imgs[imgcnt].Date, err = strTimeToUint(rec[2])
		if err != nil {
			return err
		}

		loc1, err := strconv.ParseFloat(rec[3], 32)
		if err != nil {
			return err
		}
		TagData[tagcnt].Imgs[imgcnt].Location1 = float32(loc1)

		loc2, err := strconv.ParseFloat(rec[4], 32)
		if err != nil {
			return err
		}
		TagData[tagcnt].Imgs[imgcnt].Location2 = float32(loc2)

		url1, err := strconv.ParseUint(rec[5], 10, 64)
		if err != nil {
			return err
		}
		TagData[tagcnt].Imgs[imgcnt].URL1 = uint32(url1)

		url2, err := strconv.ParseUint(rec[6], 10, 64)
		TagData[tagcnt].Imgs[imgcnt].URL2 = uint32(url2)

		TagData[tagcnt].Imgs[imgcnt].URL3 = rec[7]

		imgcnt++
	}
	runtime.GC()
	fmt.Println("csvLoader: complited!")
	return nil
}

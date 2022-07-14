package serv

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

// func ResetTags() {
// 	for _, t := range tags {
// 		for _, i := range t.id {
// 			i = "-1"
// 		}
// 	}
// }

func ReadData(dataName string) error {
	fmt.Println("csvReader: now reading data...")

	

	// read tag file
	infile, err := os.Open(dataName)
	if err != nil {
		return err
	}
	defer infile.Close()
	r := csv.NewReader(infile)
	tmp := [][]string{}
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		tmp = append(tmp, record)
	}

	tagcnt := 0
	imgcnt := 0
	for i, row := range tmp {
		if tmp[i][0] == "" {
			TagData[tagcnt].Imgs[imgcnt].Id = row[1]
			TagData[tagcnt].Imgs[imgcnt].Date = row[2]
			TagData[tagcnt].Imgs[imgcnt].Location1 = row[3]
			TagData[tagcnt].Imgs[imgcnt].Location2 = row[4]
			TagData[tagcnt].Imgs[imgcnt].URL1 = row[5]
			TagData[tagcnt].Imgs[imgcnt].URL2 = row[6]
			TagData[tagcnt].Imgs[imgcnt].URL3 = row[7]

			imgcnt++
		} else {
			if tagcnt != 0 {
				tagcnt++
			}
			imgcnt = 0

			TagData[tagcnt].Name = tmp[i][0]
			TagData[tagcnt].Imgs = make([]Image, 100)

			TagData[tagcnt].Imgs[0].Id = row[1]
			TagData[tagcnt].Imgs[0].Date = row[2]
			TagData[tagcnt].Imgs[0].Location1 = row[3]
			TagData[tagcnt].Imgs[0].Location2 = row[4]
			TagData[tagcnt].Imgs[0].URL1 = row[5]
			TagData[tagcnt].Imgs[0].URL2 = row[6]
			TagData[tagcnt].Imgs[0].URL3 = row[7]

			imgcnt++

			fmt.Println(TagData[123456].Name)
		}

	}

	fmt.Println(TagData[123456].Name)

	fmt.Println("csvReader: complited!")
	return nil
}

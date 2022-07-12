package serv

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func ReadData(tagDataName string, imgDataName string) error {
	fmt.Println("csvReader: now reading tag data...")
	// read tag file
	tagFile, err := os.Open(tagDataName)
	if err != nil {
		return err
	}
	defer tagFile.Close()
	r := csv.NewReader(tagFile)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		t := Tag{}
		t.id = record[0]
		t.name = record[1]

		tags = append(tags, t)
	}

	fmt.Println("csvReader: complited!")
	fmt.Println("csvReader: now reading img data...")

	// read img file
	imgFile, err := os.Open(imgDataName)
	if err != nil {
		return nil
	}
	defer imgFile.Close()
	r = csv.NewReader(imgFile)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		i := Image{}
		i.id = record[0]
		i.date = record[1]
		i.locate1 = record[2]
		i.locate2 = record[3]
		i.URL = record[4]

		imgs = append(imgs, i)
	}

	fmt.Println("csvReader: complited!")

	return nil
}

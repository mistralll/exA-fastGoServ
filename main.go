package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Image struct {
	Date      string
	Id        int64
	Location1 float32
	Location2 float32
	URL1      string
	URL2      string
	URL3      string
}

type Tag struct {
	Name string
	Imgs []Image
}

var TagData []Tag

func main() {
	TagData = make([]Tag, 6736297)
	err := ReadData("merged.csv")
	if err != nil {
		log.Fatal(err)
	}
	ServRun()
}

func Serch(key string) []Image {
	fmt.Println(TagData[0].Name)
	fmt.Println(key)
	ans := []Image{}
	for _, row := range TagData {
		if row.Name == key {
			ans = row.Imgs
			break
		}
	}

	fmt.Println(strconv.Itoa(len(ans)))

	return ans
}

func ReadData(dataName string) error {
	fmt.Println("csvReader: now reading data...")

	// read tag file
	infile, err := os.Open(dataName)
	if err != nil {
		return err
	}
	defer infile.Close()
	r := csv.NewReader(infile)

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
			TagData[tagcnt].Imgs = make([]Image, 100)
		}
		id, err := strconv.ParseInt(rec[1], 10, 64)
		if err != nil {
			return err
		}
		TagData[tagcnt].Imgs[imgcnt].Id = id
		TagData[tagcnt].Imgs[imgcnt].Date = rec[2]

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

		TagData[tagcnt].Imgs[imgcnt].URL1 = rec[5]
		TagData[tagcnt].Imgs[imgcnt].URL2 = rec[6]
		TagData[tagcnt].Imgs[imgcnt].URL3 = rec[7]

		imgcnt++
	}
	fmt.Println("csvReader: complited!")
	return nil
}

func ServRun() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/favicon.ico", faviconHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("tag")
	list := Serch(q)

	w.Write([]byte(`
		<html>
		<body>
	`))

	for _, row := range list {
		w.Write([]byte(`<img src=http://farm`))
		w.Write([]byte(row.URL1))
		w.Write([]byte(`.static.flickr.com/`))
		w.Write([]byte(row.URL2))
		w.Write([]byte(`/`))
		w.Write([]byte(strconv.FormatInt(row.Id, 10)))
		w.Write([]byte(`_`))
		w.Write([]byte(row.URL3))
		w.Write([]byte(`.jpg"></img>`))
		w.Write([]byte(row.Date))
		w.Write([]byte(strconv.FormatFloat(float64(row.Location1), 'f', 2, 64)))
		w.Write([]byte(strconv.FormatFloat(float64(row.Location2), 'f', 2, 64)))
	}
	w.Write([]byte(`</body></html>`))

	fmt.Println("done")
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {}

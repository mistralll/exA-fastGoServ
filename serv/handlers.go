package serv

import (
	"log"
	"net/http"
	"strconv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("tag")
	list := Serch(q)

	w.Write([]byte(`
		<html>
		<body>
	`))

	for _, row := range list {
		w.Write([]byte(`<img src=http://farm`))
		w.Write([]byte(strconv.FormatUint(uint64(row.URL1), 10)))
		w.Write([]byte(`.static.flickr.com/`))
		w.Write([]byte(strconv.FormatUint(uint64(row.URL2), 10)))
		w.Write([]byte(`/`))
		w.Write([]byte(strconv.FormatInt(row.Id, 10)))
		w.Write([]byte(`_`))
		w.Write([]byte(row.URL3))
		w.Write([]byte(`.jpg"></img>`))
		datetmp, err := uintTimeToStr(row.Date)
		if err != nil {
			log.Fatal(err)
		}
		w.Write([]byte(datetmp))
		w.Write([]byte(strconv.FormatFloat(float64(row.Location1), 'f', 2, 64)))
		w.Write([]byte(strconv.FormatFloat(float64(row.Location2), 'f', 2, 64)))
	}
	w.Write([]byte(`</body></html>`))
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {}

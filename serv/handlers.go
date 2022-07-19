package serv

import (
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("tag")
	list := Serch(q)

	w.Write([]byte(`
		<html>
		<body>
	`))

	for _, row := range list {
		tmp, err := ImageToRetImg(row)
		if err != nil {
			log.Fatal(err)
		}
		w.Write([]byte(`<img src=`))
		w.Write([]byte(tmp.url))
		w.Write([]byte(`></img>`))
		w.Write([]byte(tmp.lat))
		w.Write([]byte(tmp.lon))
		w.Write([]byte(tmp.date))
	}
	w.Write([]byte(`</body></html>`))
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {}

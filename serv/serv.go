package serv

import (
	"fmt"
	"log"
	"net/http"
)

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
		w.Write([]byte(row.Id))
		w.Write([]byte(`_`))
		w.Write([]byte(row.URL3))
		w.Write([]byte(`.jpg"></img>`))
		w.Write([]byte(row.Date))
		w.Write([]byte(row.Location1))
		w.Write([]byte(row.Location2))
	}
	w.Write([]byte(`</body></html>`))

	fmt.Println("done")
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {}

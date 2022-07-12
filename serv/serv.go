package serv

import (
	"fmt"
	"log"
	"net/http"
)

var tags []Tag
var imgs []Image

func ServRun() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/favicon.ico", faviconHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	list := Serch(r.URL.Path[1:])
	w.Write([]byte(`
		<html>
		<body>
	`))

	for _, row := range list.images {
		w.Write([]byte(`
			<img src="` + row.URL + `"></img>`))
	}
	w.Write([]byte(`</body></html>`))

	fmt.Println("done")
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {}

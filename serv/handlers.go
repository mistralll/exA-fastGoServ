package serv

import (
	"log"
	"net/http"
)

var tagAns int

func handler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("tag")

	var err error
	tagAns, err = Search(q)
	if err != nil {
		log.Fatal(err)
	}

	err = RetJson(tagAns, w)
	if err != nil {
		log.Fatal(err)
	}
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {}

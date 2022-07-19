package serv

import (
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("tag")

	ans, err := Search(q)
	if err != nil {
		log.Fatal(err)
	}

	err = RetJson(ans, w)
	if err != nil {
		log.Fatal(err)
	}
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {}

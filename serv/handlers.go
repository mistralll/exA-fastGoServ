package serv

import (
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("tag")

	ans, err := Serch(q)
	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(RetTagToHTML(ans)))
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {}

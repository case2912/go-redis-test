package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	fmt.Printf("hello golang!")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		type Ping struct {
			Status int
			Rssult string
		}
		ping := Ping{http.StatusOK, "ok"}

		res, err := json.Marshal(ping)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	})
	http.ListenAndServe(":8080", nil)
}

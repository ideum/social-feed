package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"

	"github.com/ideum/social-feed/social"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println(`Loading "/"`)

		var posts social.PostSlice = getAllPosts()

		sort.Sort(sort.Reverse(posts))

		json.NewEncoder(w).Encode(posts)
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), nil))
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"sort"

	"github.com/ideum/social-feed/social"
)

var splitter = regexp.MustCompile("[[:^word:]]*[[:word:]]+")

func main() {
	http.HandleFunc("/", socialFeedHandler)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), nil))
}

func socialFeedHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(`Loading "/"`)

	var posts social.PostSlice = getAllPosts()

	sort.Sort(sort.Reverse(posts))

	for i := range posts {
		posts[i].Text = truncate(posts[i].Text, 140)
	}

	json.NewEncoder(w).Encode(posts)
}

func truncate(s string, maxLength int) string {
	if len(s) <= maxLength {
		return s
	}

	words := splitter.FindAllString(s, -1)
	var res []byte

	for _, word := range words {
		if len(res)+len(word) > maxLength {
			break
		}
		res = append(res, word...)
	}

	return string(res) + "..."
}

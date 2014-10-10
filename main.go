package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"

	"github.com/ideum/social-feed/social"
	"github.com/ideum/social-feed/social/facebook"
	"github.com/ideum/social-feed/social/flickr"
	"github.com/ideum/social-feed/social/twitter"
	"github.com/ideum/social-feed/social/youtube"
)

var cfg struct {
	Port     int
	Twitter  twitter.Credentials
	Facebook facebook.Credentials
	Flickr   flickr.Credentials
	YouTube  youtube.Credentials
}

func init() {
	configFile, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.NewDecoder(configFile).Decode(&cfg)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Loading \"/\"")

		var posts social.PostSlice = getAllPosts()

		sort.Sort(sort.Reverse(posts))

		json.NewEncoder(w).Encode(posts)
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), nil))
}

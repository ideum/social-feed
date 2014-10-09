package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"

	"github.com/ideum/social-feed/social"
	"github.com/ideum/social-feed/social/facebook"
	"github.com/ideum/social-feed/social/twitter"
)

var cfg struct {
	Port     int `json:"-"`
	Twitter  twitter.Credentials
	Facebook facebook.Credentials
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

	flag.IntVar(&cfg.Port, "port", 8888, "HTTP Port on which to run the service")

	flag.Parse()
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Loading \"/\"")

		var posts social.PostSlice

		fApi := facebook.New(&cfg.Facebook)
		fb, _ := fApi.GetPosts()
		posts = append(posts, fb...)

		tApi := twitter.New(&cfg.Twitter)
		tweets, _ := tApi.GetPosts()
		posts = append(posts, tweets...)

		sort.Sort(sort.Reverse(posts))

		json.NewEncoder(w).Encode(posts)
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), nil))
}

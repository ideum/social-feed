package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
)

var cfg struct {
	Port    int `json:"-"`
	Twitter struct {
		ConsumerKey       string
		ConsumerSecret    string
		AccessToken       string
		AccessTokenSecret string
	}
	Facebook struct {
		AppId, AppSecret string
	}
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

		var posts SocialPostSlice

		fb, _ := GetFacebookPosts()
		posts = append(posts, fb...)

		tweets, _ := GetTwitterPosts()
		posts = append(posts, tweets...)

		sort.Sort(sort.Reverse(posts))

		json.NewEncoder(w).Encode(posts)
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), nil))
}

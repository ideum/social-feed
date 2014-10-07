package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var cfg struct {
	Port    int `json:"-"`
	Twitter struct {
		ConsumerKey       string
		ConsumerSecret    string
		AccessToken       string
		AccessTokenSecret string
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

		tweets, _ := GetTwitterPosts()
		json.NewEncoder(w).Encode(tweets)
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), nil))
}

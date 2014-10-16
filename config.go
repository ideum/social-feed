package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"

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

var cfgFile = flag.String("config", "config.json", "The JSON file containing the config")

func init() {
	flag.Parse()

	configFile, err := os.Open(*cfgFile)
	if err != nil {
		log.Fatal(err)
	}
	defer configFile.Close()

	err = json.NewDecoder(configFile).Decode(&cfg)
	if err != nil {
		log.Fatal(err)
	}
}

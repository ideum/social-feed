package youtube

import (
	"encoding/json"
	"os"
	"testing"
)

var cfg *Credentials

func init() {
	var c struct{ YouTube Credentials }
	configFile, _ := os.Open("../../config.json")
	json.NewDecoder(configFile).Decode(&c)
	cfg = &c.YouTube
}

func TestYoutubeApiConstruction(t *testing.T) {
	api := New(cfg)
	if api == nil {
		t.Fatalf("%+v", api)
	}
}

func TestYoutubeGetPosts(t *testing.T) {
	posts, err := New(cfg).GetPosts()
	if err != nil {
		t.Fatal(err)
	}
	t.Fatalf("%+v", posts)
}

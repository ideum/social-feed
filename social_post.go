package main

import (
	"encoding/json"
	"net/url"
)

type SocialPost struct {
	Text  string
	Url   url.URL
	Image url.URL
}

func (p *SocialPost) MarshalJSON() ([]byte, error) {
	tmp := struct {
		Text  string `json:"text"`
		URL   string `json:"url"`
		Image string `json:"image"`
	}{
		p.Text,
		p.Url.String(),
		p.Image.String(),
	}

	return json.Marshal(tmp)
}

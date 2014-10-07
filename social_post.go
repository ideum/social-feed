package main

import (
	"encoding/json"
	"net/url"
	"time"
)

type SocialSource string

const (
	TwitterSource SocialSource = "twitter"
)

type SocialPost struct {
	Source    SocialSource
	CreatedAt time.Time
	Text      string
	Url       url.URL
	Image     url.URL
}

func (p *SocialPost) MarshalJSON() ([]byte, error) {
	tmp := struct {
		Source    SocialSource `json:"source"`
		CreatedAt time.Time    `json:"created_at"`
		Text      string       `json:"text"`
		URL       string       `json:"url"`
		Image     string       `json:"image"`
	}{
		p.Source,
		p.CreatedAt,
		p.Text,
		p.Url.String(),
		p.Image.String(),
	}

	return json.Marshal(tmp)
}

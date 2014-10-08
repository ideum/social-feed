package main

import (
	"encoding/json"
	"net/url"
	"time"
)

type SocialSource string

const (
	TwitterSource  SocialSource = "twitter"
	FacebookSource SocialSource = "facebook"
)

type SocialPost struct {
	Source    SocialSource
	CreatedAt time.Time
	Text      string
	Url       url.URL
	Image     url.URL
}

type SocialPostSlice []SocialPost

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

func (s SocialPostSlice) Len() int           { return len(s) }
func (s SocialPostSlice) Less(i, j int) bool { return s[i].CreatedAt.Before(s[j].CreatedAt) }
func (s SocialPostSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

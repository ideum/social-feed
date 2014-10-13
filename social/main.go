package social

import (
	"encoding/json"
	"net/url"
	"time"
)

type Source string

const (
	Twitter  Source = "twitter"
	Facebook Source = "facebook"
	Flickr   Source = "flickr"
	YouTube  Source = "youtube"
)

type Post struct {
	Source    Source
	CreatedAt time.Time
	Text      string
	Url       url.URL
	Image     url.URL
}

type Provider interface {
	GetPosts() ([]Post, error)
}

type tmpPost struct {
	Source    Source    `json:"source"`
	CreatedAt time.Time `json:"created_at"`
	Text      string    `json:"text"`
	URL       string    `json:"url"`
	Image     string    `json:"image"`
}

func (p *Post) MarshalJSON() ([]byte, error) {
	tmp := tmpPost{
		p.Source,
		p.CreatedAt,
		p.Text,
		p.Url.String(),
		p.Image.String(),
	}

	return json.Marshal(tmp)
}

func (p *Post) UnmarshalJSON(j []byte) error {
	var tmp tmpPost

	err := json.Unmarshal(j, &tmp)
	if err != nil {
		return err
	}

	url, err := url.Parse(tmp.URL)
	if err != nil {
		return err
	}

	img, err := url.Parse(tmp.Image)
	if err != nil {
		return err
	}

	p.Source = tmp.Source
	p.CreatedAt = tmp.CreatedAt
	p.Text = tmp.Text
	p.Url = *url
	p.Image = *img

	return nil
}

// SocialPostSlice exists to satisfy sort.Interface.  Posts are sorted
// chronologically by their CreatedAt time.
type PostSlice []Post

func (s PostSlice) Len() int           { return len(s) }
func (s PostSlice) Less(i, j int) bool { return s[i].CreatedAt.Before(s[j].CreatedAt) }
func (s PostSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

package main

import (
	fb "github.com/huandu/facebook"
	"github.com/ideum/social-feed/social"
	"net/url"
	"time"
)

func GetFacebookPosts() ([]social.Post, error) {
	app := fb.New(cfg.Facebook.AppId, cfg.Facebook.AppSecret)
	token := app.AppAccessToken()
	session := app.Session(token)

	res, _ := session.Get("/242598822781/feed", fb.Params{})

	var data []struct {
		Id          string
		CreatedTime string
		Message     string
		Picture     string
	}

	err := res.DecodeField("data", &data)

	if err != nil {
		return nil, err
	}

	posts := make([]social.Post, 0, len(data))

	// all facebook pages should link back to the master timeline
	u, _ := url.Parse("http://facebook.com/ideum")

	for _, d := range data {
		// skip fluff content like likes and friend requests
		if d.Message == "" && d.Picture == "" {
			continue
		}

		t, _ := time.Parse("2006-01-02T15:04:05-0700", d.CreatedTime)
		i, _ := url.Parse(d.Picture)

		posts = append(posts, social.Post{
			Source:    social.Facebook,
			CreatedAt: t,
			Text:      d.Message,
			Url:       *u,
			Image:     *i,
		})
	}

	return posts, nil
}

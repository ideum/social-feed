package main

import (
	"github.com/ideum/social-feed/social"
	"github.com/ideum/social-feed/social/facebook"
	"github.com/ideum/social-feed/social/flickr"
	"github.com/ideum/social-feed/social/twitter"
	"github.com/ideum/social-feed/social/youtube"
)

func getAllPosts() []social.Post {
	posts := []social.Post{}
	pc := make(chan []social.Post)

	providers := [...]social.Provider{
		twitter.New(&cfg.Twitter),
		facebook.New(&cfg.Facebook),
		flickr.New(&cfg.Flickr),
		youtube.New(&cfg.YouTube),
	}

	for _, p := range providers {
		go func(p social.Provider) {
			pposts, _ := p.GetPosts()
			pc <- pposts
		}(p)
	}

	for _ = range providers {
		posts = append(posts, <-pc...)
	}

	return posts
}

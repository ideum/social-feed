package main

import (
	"github.com/ideum/social-feed/social"
	"github.com/ideum/social-feed/social/facebook"
	"github.com/ideum/social-feed/social/twitter"
)

func getAllPosts() []social.Post {
	posts := []social.Post{}
	pc := make(chan []social.Post)

	providers := [...]social.Provider{
		facebook.New(&cfg.Facebook),
		twitter.New(&cfg.Twitter),
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

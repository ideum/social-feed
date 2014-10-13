package main

import (
	"github.com/ideum/social-feed/social"
)

func getPosts(p social.Provider, s social.Source, pc chan social.Post) {
	posts, err := p.GetPosts()
	if err != nil {
		return
	}

	for _, post := range posts {
		pc <- post
	}
}

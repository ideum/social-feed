package main

import (
	"sync"

	"github.com/ideum/social-feed/social"
	"github.com/ideum/social-feed/social/facebook"
	"github.com/ideum/social-feed/social/flickr"
	"github.com/ideum/social-feed/social/twitter"
	"github.com/ideum/social-feed/social/youtube"
)

func getAllPosts() []social.Post {
	var wg sync.WaitGroup
	posts := []social.Post{}
	pc := make(chan social.Post)

	providers := map[social.Source]social.Provider{
		social.Twitter:  twitter.New(&cfg.Twitter),
		social.Facebook: facebook.New(&cfg.Facebook),
		social.Flickr:   flickr.New(&cfg.Flickr),
		social.YouTube:  youtube.New(&cfg.YouTube),
	}

	wg.Add(len(providers))

	for s, p := range providers {
		go func(s social.Source, p social.Provider) {
			getPosts(p, s, pc)
			wg.Done()
		}(s, p)
	}

	go func() {
		wg.Wait()
		close(pc)
	}()

	for post := range pc {
		posts = append(posts, post)
	}

	return posts
}

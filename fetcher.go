package main

import (
	"encoding/json"
	"math/rand"

	"github.com/bradfitz/gomemcache/memcache"

	"github.com/ideum/social-feed/social"
)

var mc = memcache.New("127.0.0.1:11211")

func getPosts(p social.Provider, s social.Source, pc chan social.Post) {
	var posts []social.Post
	key := "social-posts-" + string(s)

	if cache, err := mc.Get(key); err == nil {
		// Cache hit; use cached version

		err = json.Unmarshal(cache.Value, &posts)
		if err != nil {
			return
		}
	} else {
		// Cache miss; fetch posts over wire

		posts, err = p.GetPosts()
		if err != nil {
			return
		}

		// random expiration time between 1 and 2 hours
		expiration := 3600 + rand.Int31n(3600)

		value, err := json.Marshal(posts)
		if err == nil {
			mc.Set(&memcache.Item{
				Key:        key,
				Value:      value,
				Expiration: expiration,
			})
		}
	}

	for _, post := range posts {
		pc <- post
	}
}

package main

import "math/rand/v2"

var id uint = 1

type Post struct {
	Title, Contents, Author string
}

func assignID() uint {
	id++
	return id
}

type blog struct {
	Posts map[uint]Post
}

func Newassign(blog *blog) uint {
	for {
		id := uint(rand.IntN(1000))
		if _, exists := blog.Posts[id]; !exists {
			return id
		}
	}
}

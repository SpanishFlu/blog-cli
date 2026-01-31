package main

import "math/rand/v2"

type Post struct {
	Title, Contents, Author string
	ID                      uint
}

func assignID(blog *blog) uint {
	for {
		id := uint(rand.IntN(1000))
		if !checkID(blog, id) {
			return id
		}
	}
}

type blog struct {
	Posts []Post
}

func checkID(blog *blog, id uint) bool {
	for _, post := range blog.Posts {
		if post.ID == id {
			return true
		}
	}
	return false
}

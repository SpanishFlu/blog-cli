package main

import "math/rand/v2"

var id uint = 0

type Post struct {
	Title, Contents, Author string
}

func assignID(blog *blog) uint {
	id++
	return id
}

type blog struct {
	Posts map[uint]Post
}

/*func checkID(blog *blog, id uint) bool {
	for _, post := range blog.Posts {
		if post.ID == id {
			return true
		}
	}
	return false
}*/

func Newassign(blog *blog) uint {
	for {
		id := uint(rand.IntN(1000))
		if _, exists := blog.Posts[id]; !exists {
			return id
		}
	}
}

package main

//import "math/rand/v2"

type Post struct {
	Title, Contents, Author string
	ID                      uint
}

func assignID(blog *blog) uint {
	var pid uint
	pid = 0
	pid++
	return pid
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

/*func Newassign(blog *blog) uint{
	ID := make(map[uint]uint)
	for
	ID[]
}*/

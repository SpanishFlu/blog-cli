package main

import (
	"fmt"
	"strings"
)

func creatingPost(blog *blog) {
	var title, contents, author string
	fmt.Println("Please Enter The Title:")
	fmt.Scanln(&title)
	if err := validationCreatedPost(title); err != nil {
		fmt.Println("Post creation failed:", err)
		return
	}

	fmt.Println("Please Type Your Post:")
	fmt.Scanln(&contents)
	if err := validationCreatedPost(contents); err != nil {
		fmt.Println("Post creation failed:", err)
		return
	}

	fmt.Println("Please Enter The Author Name:")
	fmt.Scanln(&author)
	author = strings.TrimSpace(author)
	if err := validationCreatedPost(author); err != nil {
		fmt.Println("Post creation failed:", err)
		return
	}

	x := savingPost(blog, title, contents, author)
	fmt.Println(">>> Post Created Successfully <<<")
	fmt.Printf("your Post ID is: %d\n", x)
}

func savingPost(blog *blog, title, contents, author string) uint {
	id := Newassign(blog)
	blog.Posts[id] = Post{
		Title:    title,
		Contents: contents,
		Author:   author,
	}
	// same as p.Title = title
	//blog.Posts = append(blog.Posts, blog.Posts[id])
	return id
}

func viewingPosts(blog *blog) {
	if len(blog.Posts) == 0 {
		fmt.Println("No posts available.")
		return
	}
	for id, p := range blog.Posts {
		fmt.Println("---------------")
		fmt.Printf("Title: %s\n", p.Title)
		fmt.Printf("%s\n", p.Contents)
		fmt.Printf("Author: %s\n", p.Author)
		fmt.Printf("ID: %d\n", id)
	}
}

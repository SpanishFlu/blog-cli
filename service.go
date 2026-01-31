package main

import "fmt"

func creatingPost() (string, string, string) {
	var title, contents, author string

	fmt.Println("Please Enter The Title")
	fmt.Scanln(&title)

	fmt.Println("Please Type Your Post")
	fmt.Scanln(&contents)

	fmt.Println("Please Enter The Author Name")
	fmt.Scanln(&author)
	validationCreatedPost(title, contents, author)
	fmt.Println("Post Created Successfully")
	return title, contents, author
}

func savingPost() {
	var Blog blog
	title, contents, author := creatingPost()
	p := Post{
		Title:    title,
		Contents: contents,
		Author:   author,
		ID:       assignID(&Blog),
	}
	// same as p.Title = title
	Blog.Posts = append(Blog.Posts, p)
}

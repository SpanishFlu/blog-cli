package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func creatingPost(blog *blog) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please Enter The Title:")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Println("Please Type Your Post:")
	contents, _ := reader.ReadString('\n')
	contents = strings.TrimSpace(contents)

	fmt.Println("Please Enter The Author Name:")
	author, _ := reader.ReadString('\n')
	author = strings.TrimSpace(author)

	validationCreatedPost(title, contents, author)
	savingPost(blog, title, contents, author)
	fmt.Println(">>> Post Created Successfully <<<")
}

func savingPost(blog *blog, title, contents, author string) {
	p := Post{
		Title:    title,
		Contents: contents,
		Author:   author,
		ID:       assignID(blog),
	}
	// same as p.Title = title
	blog.Posts = append(blog.Posts, p)
}

func viewingPosts(blog *blog) {
	if len(blog.Posts) == 0 {
		fmt.Println("No posts available.")
		return
	}
	for _, p := range blog.Posts {
		fmt.Println("---------------")
		fmt.Printf("Title: %s\n", p.Title)
		fmt.Printf("%s\n", p.Contents)
		fmt.Printf("Author: %s\n", p.Author)
	}
}

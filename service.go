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
	if err := validationCreatedPost(title); err != nil {
		fmt.Println("Post creation failed:", err)
		return
	}

	fmt.Println("Please Type Your Post:")
	contents, _ := reader.ReadString('\n')
	contents = strings.TrimSpace(contents)
	if err := validationCreatedPost(contents); err != nil {
		fmt.Println("Post creation failed:", err)
		return
	}

	fmt.Println("Please Enter The Author Name:")
	author, _ := reader.ReadString('\n')
	author = strings.TrimSpace(author)
	if err := validationCreatedPost(author); err != nil {
		fmt.Println("Post creation failed:", err)
		return
	}

	x := savingPost(blog, title, contents, author)
	fmt.Println(">>> Post Created Successfully <<<")
	fmt.Printf("your Post ID is: %d\n", x)
}

func savingPost(blog *blog, title, contents, author string) int {
	p := Post{
		Title:    title,
		Contents: contents,
		Author:   author,
		ID:       assignID(blog),
	}
	// same as p.Title = title
	blog.Posts = append(blog.Posts, p)
	return int(p.ID)
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
		fmt.Printf("ID: %d\n", p.ID)
	}
}

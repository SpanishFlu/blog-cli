package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var Blog blog

/*
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
*/
func creatingPost() (string, string, string) {
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
	fmt.Println(">>> Post Created Successfully <<<")

	return title, contents, author
}

func savingPost() {
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

/*func ViewPostCreated() {
	title, contents, author := creatingPost()
	fmt.Printf("Title: %s", title)
	fmt.Printf("Contents: %s", contents)
	fmt.Printf("Author: %s", author)
}*/

func viewingPosts(blog *blog) {
	for _, post := range blog.Posts {
		fmt.Println("---------------")
		fmt.Printf("Title: %s\n", post.Title)
		fmt.Printf("%s\n", post.Contents)
		fmt.Printf("Author: %s\n", post.Author)
	}
}

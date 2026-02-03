package main

import (
	"fmt"
)

func creatingPost(blog *blog) {
	var title, contents, author string
	fmt.Println("Please Enter The Title:")
	fmt.Scanln(&title)
	if err := validationCreatedPost(title); err != nil {
		fmt.Println(err, "Title can't be empty")
		return
	}

	fmt.Println("Please Type Your Post:")
	fmt.Scanln(&contents)
	if err := validationCreatedPost(contents); err != nil {
		fmt.Println(err, "No contents have been entered")
		return
	}

	fmt.Println("Please Enter The Author Name:")
	fmt.Scanln(&author)
	if err := validationCreatedPost(author); err != nil {
		fmt.Println(err, "The auther name is empty")
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

func ViewPostbyId(blog *blog) {
	var id uint
	if err := checkEmpty(blog); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print("Enter Post ID: ")
	fmt.Scan(&id)
	post, _ := blog.Posts[id]

	if checkIfPostExists(blog, id) {
		fmt.Println("Title:", post.Title)
		fmt.Println("Contents:", post.Contents)
		fmt.Println("Author:", post.Author)
	} else {
		fmt.Println("No post with this ID")
		return
	}

}

func updatePost(blog *blog) {
	var title, contents, author string
	fmt.Printf("please enter post ID")
	fmt.Scanln(&id)
	post := blog.Posts[id]
	fmt.Printf("Title: %s\n", post.Title)
	fmt.Printf("%s\n", post.Contents)
	fmt.Printf("Author: %s\n", post.Author)

	fmt.Println("New Title: ")
	fmt.Scanln(&title)
	if title == post.Title {
		fmt.Println("The Title is the same...")
	}

	fmt.Println("Edit Contents: ")
	fmt.Scanln(&contents)
	if contents == post.Contents {
		fmt.Println("The Contents are the same...")
	}

	fmt.Println("Edit Author Name: ")
	fmt.Scanln(&author)
	if author == post.Author {
		fmt.Println("The Author is the same...")
	}
}

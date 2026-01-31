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
	p := post{}
	title, contents, author := creatingPost()
	p.title = append(p.title, title)
	p.contents = append(p.contents, contents)
	p.author = append(p.author, author)

}

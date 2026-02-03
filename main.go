package main

import "fmt"

func main() {
	var Choose int
	fmt.Println("Blog CLI start")
	myBlog := &blog{
		Posts: make(map[uint]Post),
	}

	for {
		printmenu()
		fmt.Scanln(&Choose)
		switch Choose {
		case 1:
			fmt.Println("Create Post selected")
			creatingPost(myBlog)
		case 2:
			fmt.Println("Viewing All Posts")
			viewingPosts(myBlog)
		case 3:
			fmt.Println("View Post by ID selected")
			checkEmpty(myBlog)
			ViewPostbyId(myBlog)
			return
		case 4:
			fmt.Println("Update Post Selected")
		case 5:
			fmt.Println("Delete Post selected")
		case 6:
			fmt.Println("Search Posts Selected")
		case 0:
			fmt.Println("Alright, waiting for the next blog (;")
			return
		default:
			fmt.Println("Invalid Choice Please Try Again")
		}
	}
}

func printmenu() {
	fmt.Println("1. Create Post")
	fmt.Println("2. View All Posts")
	fmt.Println("3. View Post by ID")
	fmt.Println("4. Update Post")
	fmt.Println("5. Delete Post")
	fmt.Println("6. Search Posts")
	fmt.Println("0. Exit")
	fmt.Println("Choose an option:")
}

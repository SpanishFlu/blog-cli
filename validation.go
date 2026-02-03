package main

import (
	"errors"
	"fmt"
)

func validationCreatedPost(x string) error {
	if x == "" {
		return errors.New("this part can't be empty")
	}
	return nil
}

func checkIfPostExists(blog *blog, id uint) bool {
	_, exists := blog.Posts[id]
	if len(blog.Posts) == 0 {
		fmt.Println("No posts has been created yet xoxo")
	}
	return exists
}

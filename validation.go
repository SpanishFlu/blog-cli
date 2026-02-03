package main

import (
	"errors"
)

func validationCreatedPost(x string) error {
	if x == "" {
		return errors.New("Post creation failed: ")
	}
	return nil
}

func checkIfPostExists(blog *blog, id uint) bool {
	_, exists := blog.Posts[id]
	return exists
}

func checkEmpty(blog *blog) error {
	if len(blog.Posts) == 0 {
		return errors.New("No posts has been created yet xoxo")
	}
	return nil
}

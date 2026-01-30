package main

import (
	"errors"
)

func validationCreatedPost(title, contents, author string) error {
	if title == "" {
		errors.New("Title is empty")
	} else if contents == "" {
		errors.New("No contents were entered")
	} else if author == "" {
		errors.New("Author Name wasn't entered")
	}
	return nil
}

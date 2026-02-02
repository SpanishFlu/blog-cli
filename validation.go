package main

import (
	"errors"
)

func validationCreatedPost(title, contents, author string) error {
	if title == "" {
		return errors.New("Title is empty")
	}
	if contents == "" {
		return errors.New("No contents were entered")
	}
	if author == "" {
		return errors.New("Author Name wasn't entered")
	}
	return nil
}

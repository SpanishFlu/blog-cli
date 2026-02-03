package main

import (
	"errors"
)

func validationCreatedPost(x string) error {
	if x == "" {
		return errors.New("this part can't be empty")
	}
	return nil
}

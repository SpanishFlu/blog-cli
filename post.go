package main

import "math/rand/v2"

type post struct {
	title, contents, author []string
	ID                      []int
}

func assignID(p *post) {
	p.ID = append(p.ID, rand.IntN(100))
}

package main

import (
	"log"

	"github.com/haunt98/changeloguru/internal/git"
)

func main() {
	r, err := git.NewRepository(".")
	if err != nil {
		log.Fatal(err)
	}

	commits, err := r.Log("", "")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Commits: ", commits)

	tags, err := r.SemVerTags()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Tags: ", tags)
}

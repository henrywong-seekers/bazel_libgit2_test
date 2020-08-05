package main

import (
	"log"

	"github.com/libgit2/git2go/v30"
)

func main() {
	_, err := git.Clone("https://github.com/whs-dot-hk/ansible-fedora-32.git", "test", &git.CloneOptions{})
	if err != nil {
		log.Fatal(err)
	}
}

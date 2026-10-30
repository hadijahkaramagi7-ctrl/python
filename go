package main

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/v60/github"
	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")})
	client := github.NewClient(oauth2.NewClient(ctx, ts))

	content := []byte("print('hello')\n")
	opts := &github.RepositoryContentFileOptions{
		Message: github.String("Add file"),
		Content: content,
		Branch:  github.String("main"),
	}
	_, _, err := client.Repositories.CreateFile(ctx, "USER", "REPO", "hello.py", opts)
	if err != nil { panic(err) }
	fmt.Println("✅ Uploaded")
}

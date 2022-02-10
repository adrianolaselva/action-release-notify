package main

import (
	"context"
	"fmt"
	"syscall"

	"github.com/google/go-github/v42/github"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

func main() {
	_ = godotenv.Load(".env")

	accessToken := ""
	if value, ok := syscall.Getenv("API_TOKEN_GITHUB"); ok {
		accessToken = value
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	tags, _, err := client.Repositories.ListTags(ctx, "paygoc6", "transaction-paygoweb-plataforma", nil)
	if err != nil {
		panic(err)
	}

	for k, v := range tags {
		fmt.Printf("%v => %v\n", k, *v.Name)

		tag, _, _ := client.Repositories.GetReleaseByTag(ctx, "paygoc6", "transaction-paygoweb-plataforma", *v.Name)

		fmt.Printf("%v\n", *tag.ID)
		fmt.Printf("%v\n", *tag.AssetsURL)
		fmt.Printf("%v\n", *tag.Author.Login)
		fmt.Printf("%v\n", *tag.URL)
		fmt.Printf("%v\n", *tag.PublishedAt)
		fmt.Printf("%v\n", *tag.TagName)
		fmt.Printf("%v\n", *tag.TarballURL)
		fmt.Printf("%v\n", *tag.Body)
		break
	}
}

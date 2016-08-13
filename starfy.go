package main

import (
	"fmt"
	"log"
	"os"

	"github.com/chck/starfy/config"
	"github.com/google/go-github/github"
	flags "github.com/jessevdk/go-flags"
	"golang.org/x/oauth2"
)

type options struct {
	Verbose []bool `short:"v" long:"verbose" description:"Show verbose debug information"`
}

func args() (string, string) {
	opts := &options{}
	parser := flags.NewParser(opts, flags.PrintErrors)
	parser.Usage = "[OWNER] [REPO]"
	args, _ := parser.Parse()

	if len(args) < 2 {
		parser.WriteHelp(os.Stdout)
		os.Exit(1)
	}

	return args[0], args[1]
}

func star(owner string, repo string, token string, errCh chan error) {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	client := github.NewClient(tc)

	res, err := client.Activity.Star(owner, repo)
	log.Println(res)

	if err != nil {
		errCh <- err
	} else {
		errCh <- nil
	}
}

func main() {
	owner, repo := args()
	tokens := config.Load().Github.Tokens

	errCh := make(chan error, 1)

	stars := 0
	for _, token := range tokens {
		go star(owner, repo, token, errCh)
		err := <-errCh
		if err == nil {
			stars++
		}
	}

	fmt.Println(stars)
}

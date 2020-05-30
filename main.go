package main

import (
	"github.com/TaigaMikami/go-dev/api"
	"github.com/k0kubun/pp"
)

func main() {
	opt := &api.RetrieveArticlesOption{
		Page:    1,
		PerPage: 1,
		State:   api.StateFresh,
	}

	client := api.NewClient("")
	res, err := client.RetrieveArticles(opt)
	if err != nil {
		panic(err)
	}
	pp.Print(res)
}

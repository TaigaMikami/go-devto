package main

import (
	"github.com/TaigaMikami/go-devto/api"
	"github.com/k0kubun/pp"
)

func main() {
	opt := &api.RetrieveVideoArticlesOption{
		Page:    1,
		PerPage: 1,
	}

	client := api.NewClient("")
	res, err := client.RetrieveVideoArticles(opt)
	if err != nil {
		panic(err)
	}
	pp.Print(res)
}

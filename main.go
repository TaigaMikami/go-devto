package main

import (
	"github.com/TaigaMikami/go-devto/api"
	"github.com/k0kubun/pp"
)

func main() {
	opt := &api.RetrieveFollowersOption{
		Page:    1,
		PerPage: 1,
	}

	client := api.NewClient("")
	res, err := client.RetrieveFollowers(opt)
	if err != nil {
		panic(err)
	}
	pp.Print(res)
}

//func main() {
//	opt := &api.RetrieveCommentsOption{
//		AId: 270180,
//	}
//
//	client := api.NewClient("")
//	res, err := client.RetrieveComments(opt)
//	if err != nil {
//		panic(err)
//	}
//	pp.Print(res)
//}

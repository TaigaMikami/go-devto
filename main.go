package main

import (
	"github.com/TaigaMikami/go-devto/api"
	"github.com/k0kubun/pp"
)

//func main() {
//	opt := &api.RetrieveFollowersOption{
//		Page:    1,
//		PerPage: 1,
//	}
//
//	client := api.NewClient("")
//	res, err := client.RetrieveFollowers(opt)
//	if err != nil {
//		panic(err)
//	}
//	pp.Print(res)
//}

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

//func main() {
//	opt := &api.DraftArticle{
//		Article: api.ArticleContent{
//			Title:        "Hello, World!",
//			Published:    false,
//			BodyMarkdown: "Hello DEV, this is my first post",
//			Tags:         []string{"discuss", "javascript"},
//		},
//	}
//
//	client := api.NewClient("")
//	res, err := client.AddArticle(opt)
//	if err != nil {
//		panic(err)
//	}
//
//	pp.Print(res)
//}

//func main() {
//	opt := &api.RetrieveUserArticlesOption{
//		Page:    1,
//		PerPage: 2,
//	}
//
//	client := api.NewClient("BvGbopywQrMEFaiNHsNFLsGZ")
//	res, _ := client.RetrieveUserAllArticles(opt)
//	pp.Print(res)
//}

func main() {
	opt := &api.DraftArticle{
		Article: api.ArticleContent{
			Title:        "Hello, World! modify",
			Published:    false,
			BodyMarkdown: "Hello DEV, this is my first post modify",
			Tags:         []string{"discuss", "javascript"},
		},
	}

	client := api.NewClient("")
	res, err := client.ModifyArticle(346946, opt)
	if err != nil {
		panic(err)
	}

	pp.Print(res)
}

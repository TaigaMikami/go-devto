<p align="center">
  <img src="https://user-images.githubusercontent.com/25325947/83355285-723ac780-a399-11ea-9413-2689f86671de.png" />
</p>

# go-devto
![Test](https://github.com/TaigaMikami/go-devto/workflows/Test/badge.svg)
![License](https://img.shields.io/github/license/TaigaMikami/go-devto)

This is a Go wrapper for working with [DEV API](https://docs.dev.to/api/).

The aim is to be able to do all tasks on this DEV API endpoint simply by getting an API key.

## Install
```
go get github.com/TaigaMikami/go-devto/devto
```

## Usage
1. Visit https://dev.to/settings/account and Get API key.

2. eg.) Retrieve articles and Output an article's title.

```go
package main

import (
	"fmt"
	"github.com/TaigaMikami/go-devto/devto"
)

func main() {
	opt := &devto.RetrieveArticlesOption{
		Page: 1,
		PerPage: 10,
	}
	client := devto.NewClient("API Key")
	res, err := client.RetrieveArticles(opt)
	if err != nil {
		panic(err)
	}

	fmt.Println(res[0].Title)
}
```

## API Examples
Examples of the API can be found in the [examples](https://github.com/TaigaMikami/go-devto/tree/master/example) directory.


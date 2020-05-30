package api

import (
	"encoding/json"
	"fmt"
	"github.com/TaigaMikami/go-devto/util"
	"github.com/k0kubun/pp"
	"net/http"
	"reflect"
	"strconv"
)

var Origin = "https://dev.to/api"
var DefaultClient = http.DefaultClient

type Client struct {
	ApiKey string
}

func NewClient(apiKey string) *Client {
	return &Client{
		ApiKey: apiKey,
	}
}

func getJSON(req *http.Request, res interface{}) error {
	resp, err := DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		pp.Println(err)
	}

	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(res)
}

func GetQuery(action string, data, res interface{}) error {
	req, err := http.NewRequest("GET", Origin+action, nil)
	if err != nil {
		return err
	}

	// generate query parameter
	params := req.URL.Query()
	v := reflect.Indirect(reflect.ValueOf(data))
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		key := util.ToSnakeCase(t.Field(i).Name)
		f := v.Field(i).Interface()
		if val, ok := f.(int); ok {
			if val != 0 {
				params.Add(key, strconv.Itoa(val))
			}
		} else {
			if len(v.Field(i).String()) != 0 {
				params.Add(key, v.Field(i).String())
			}
		}
	}
	req.URL.RawQuery = params.Encode()
	fmt.Println(req.URL.String())
	return getJSON(req, res)
}

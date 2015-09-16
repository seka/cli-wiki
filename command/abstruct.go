package command

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/codegangsta/cli"
)

const (
	url = "http://ja.wikipedia.org/w/api.php?action=query&format=json&&prop=extracts&redirects=1&exchars=130&explaintext=1"
)

// Param api parameters
type Param struct {
	Titles string
}

// Contents API Response
type Contents struct {
	Extract string  `json:"extract"`
	Pageid  float64 `json:"pageid"`
	Title   string  `json:"title"`
}

// CmdAbstruct ...
func CmdAbstruct(c *cli.Context) {
	var param Param
	if args := c.Args(); len(args) > 0 {
		param.Titles = "titles=" + args[0]
	}

	request := url + "&" + param.Titles
	response, err := http.Get(request)
	defer response.Body.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	byt, _ := ioutil.ReadAll(response.Body)

	var resAll map[string]interface{}
	if err := json.Unmarshal(byt, &resAll); err != nil {
		fmt.Println(err)
		return
	}
	resQuery := resAll["query"].(map[string]interface{})
	resPages := resQuery["pages"].(map[string]interface{})

	var contents Contents
	for key, val := range resPages {
		if key == "-1" {
			fmt.Println("not found...")
			return
		}
		m := val.(map[string]interface{})
		contents = Contents{
			Extract: m["extract"].(string),
			Pageid:  m["pageid"].(float64),
			Title:   m["title"].(string),
		}
	}

	fmt.Println("--------------------------------------")
	fmt.Println(contents.Title)
	fmt.Println("--------------------------------------")
	fmt.Println(contents.Extract)
	fmt.Println("--------------------------------------")

}

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var Stacks = map[string]string{
	"math":       "math.stackexchange",
	"physics":    "physics.stackexchange",
	"overflow":   "stackoverflow",
	"scifi":	  "scifi.stackexchange",
	"software":   "software.stackexchange",
	"everywhere": "stackoverflow",
}

var StackExchangeClient Client

func init() {
	// Initialise client attributes
	StackExchangeClient.Host = "api.stackexchange.com"
	StackExchangeClient.BasePath = "2.2"
	StackExchangeClient.Protocol = "https"
	StackExchangeClient.Filter = "!1PVL)ZRqiPq7_k1QEI)oXUiKFkvaVNDWT"
}

type Model struct {
	Verb, Adjective, Query, Stack string
}

type Client struct {
	Host, BasePath, Protocol, Filter string
	Model                            Model
}

func (c *Client) constructSearchRequest() string {

	var query string
	var url string

	if strings.Contains(c.Model.Query, " ") {
		query = strings.Replace(c.Model.Query, " ", "%", -1)
	} else {
		query = c.Model.Query
	}

	parameters := fmt.Sprintf(
		"?page=1&pagesize=1&order=desc&sort=creation&intitle=%s&site=%s",
		query, Stacks[c.Model.Stack])

	request := fmt.Sprintf(
		"%s://%s/%s/search",
		c.Protocol, c.Host, c.BasePath)

	url = request + parameters
	return url
}

func (c *Client) GetSearchRequest() (SearchAPIResponse, error) {
	url := c.constructSearchRequest()
	s := SearchAPIResponse{}

	response, respErr := http.Get(url)
	if respErr != nil {
		return s, respErr
	} else {
		defer response.Body.Close()
		contents, readErr := ioutil.ReadAll(response.Body)
		if readErr != nil {
			return s, readErr
		}
		parseErr := json.Unmarshal(contents, &s)
		if parseErr != nil {
			return s, parseErr
		}
		return s, nil
	}
}

func (c *Client) Query() {
	if c.Model.Query != "" {
		if c.Model.Adjective != "" {

			response, _ := c.GetSearchRequest()

			if len(response.Items) == 0 {
				fmt.Println("No matches found")
			} else {
				fmt.Printf(
					"Matched: %s \n Do you want to browse? (y/n) \n",
					response.Items[0].Title,
				)
				// Handle user selection for how to proceed
				scanner := bufio.NewScanner(os.Stdin)
				scanner.Scan()

				if scanner.Err() != nil {
					fmt.Printf("Error in decision: %s", scanner.Err())
				}

				decision := scanner.Text()
				fmt.Printf("You selected: %s \n", decision)
			}

		} else {
			fmt.Println("Must select query adjective!")
		}

	} else {
		fmt.Println("You must enter a query!")
	}
}


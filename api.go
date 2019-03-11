package main

import (
	"bufio"
	"fmt"
	"os"
)

// Query
func Query(verb, adjective, query, stack string) {
	if query != "" {
		if adjective != "" {

			c := StackExchangeClient
			c.Model.Verb = verb
			c.Model.Adjective = adjective

			c.Model.Stack = stack
			c.Model.Query = query

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


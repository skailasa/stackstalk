package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
)

func query(verb, adjective, query, stack string) {
	if query != "" {

		if adjective != "" {
			c := Client{
				Host:     "",
				BasePath: "",
				Protocol: "",
				Model:    Model{
					verb,
					adjective,
					query,
					stack,
				},
			}

			c.GetRequest()

			fmt.Printf(
				"Searching %s for the %s query, matching `%s` \n",
				stack, adjective, query,
			)
		} else {
			fmt.Println("must select query adjective!")
		}

	} else {
		fmt.Println("You must enter a query!")
	}
}

func subCommandFactory(name string) cli.Command {
		return cli.Command{
			Name: name,
			Action: func(c *cli.Context) error {
				fmt.Printf("Returning %s result", name)
				query(
					"query",
					name,
					c.Args().Get(0),
					c.String("stack"),
				)
				return nil
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "stack, s",
					Value: "everywhere",
					Usage: "Select a `STACK` to stalk, e.g. 'math''",
				},
			},
    	}
}


func main() {
	app := cli.NewApp()
	app.Author = "Srinath Kailasa"
	app.Email = "srinathkailasa@gmail.com"

	app.Name = "Stack Stalk"
	app.Version = "0.0.0"
	app.Usage = "stalk the stack universe"
	app.Description = "Tired of endless stack X browser tabs? Here's a nifty \n" +
		"   CLI instead! I grew sick of endless tabs draining my \n" +
		"   battery and cluttering my browser. So I decided to \n" +
		"   take a stand on building the command line app I \n" +
		"   always dreamed of."

	// Top level commands are <verbs> e.g. 'query'
	// Second level commands are <adjectives> e.g. 'new'

	app.Flags = []cli.Flag {
			cli.StringFlag{
				Name: "stack, s",
				Value: "everywhere",
				Usage: "Select a `STACK` to stalk, e.g. 'math''",
			},
	}

	app.Commands  = []cli.Command {
		cli.Command{
			Name: "query",
			Before: func(c *cli.Context) error {
				fmt.Println("Searching the stack universe!")
				return nil
			},
			Action: func(c *cli.Context) error {
				query(
					"query",
					"" ,
					c.Args().Get(0),
					c.String("stack"),
					)
				return nil
			},
			Subcommands: []cli.Command {
				subCommandFactory("top"),
				subCommandFactory("new"),
				subCommandFactory("old"),
				subCommandFactory("wildcard"),
			},
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
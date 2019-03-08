package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
)

func query(query, stack string) {
	if query != "" {
		fmt.Printf("Searching %s for `%s` \n", stack, query)
	} else {
		fmt.Println("You must enter a query!")
	}
}

func main() {
	app := cli.NewApp()
	app.Author = "Srinath Kailasa"
	app.Email = "srinathkailasa@gmail.com"

	app.Name = "Stack Stalk"
	app.Version = "0.0.0"
	app.Usage = "stalk the stack universe"
	app.Description = "Tired of endless stack X browser tabs? Here's a nifty CLI instead! \n" +
					  "   I grew sick of endless tabs draining my battery and cluttering my \n" +
					  "   browser. So I decided to take a stand on building the command line \n" +
					  "   app I always dreamed of."

	app.Commands  = []cli.Command {
		cli.Command{
			Name: "query",
			Before: func(c *cli.Context) error {
				fmt.Println("Searching the stack universe!")
				return nil
			},
			Action: func(c *cli.Context) error{
				query(c.Args().Get(0), c.String("stack"))
				return nil
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "stack, s",
					Value: "everywhere",
					Usage: "Select a `STACK` to stalk, e.g. 'math''",
				},
			},
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
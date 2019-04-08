package main

import (
	"BtcdTest/core"
	"bufio"
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
	"strings"
)

func main() {
	app := cli.NewApp()


	app.Action = func(c *cli.Context) {
		if c.NArg() != 0 {
			fmt.Printf("未找到命令: %s\n运行命令 %s help 获取帮助\n", c.Args().Get(0), app.Name)
			return
		}

		var prompt  string

		prompt = app.Name + " > "
	L:
		for {
			var input string
			fmt.Print(prompt)
			//   fmt.Scanln(&input)

			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan() // use `for scanner.Scan()` to keep reading
			input = scanner.Text()
			//fmt.Println("captured:",input)
			switch input {
			case "close":
				fmt.Println("close.")
				break L
			default:
			}
			//fmt.Print(input)
			cmdArgs := strings.Split(input, " ")
			//fmt.Print(len(cmdArgs))
			if len(cmdArgs) == 0 {
				continue
			}

			s := []string{app.Name}
			s = append(s,cmdArgs...)

			c.App.Run(s)

		}

		return
	}

	app.Commands = []cli.Command{
		{
			Name:    "blockchain",
			Aliases: []string{"bc"},
			Usage:   "show the blockchain",
			Action:  func(c *cli.Context)error{
				core.ShowBlockchain(c)
				return nil
			},
		},
		{
			Name:    "block",
			Aliases: []string{"b"},
			Usage:   "complete a task on the list",
			Action:  func(c *cli.Context) error {
				core.ShowBlockByNumber(c)
				return nil
			},
		},
		{
			Name:        "create",
			Aliases:     []string{"c"},
			Usage:       "options for create",
			Subcommands: []cli.Command{
				{
					Name:  "Tx",
					Usage: "create a transaction",
					Aliases:     []string{"tx"},
					Action: func(c *cli.Context) error {
						fmt.Printf("A send %s to B\n",c.Args().First())
						return nil
					},
				},
				{
					Name:  "Block",
					Usage: "create a block",
					Aliases:     []string{"b"},
					Action: func(c *cli.Context) error {
						fmt.Println("create a block !")
						return nil
					},
				},
			},
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}


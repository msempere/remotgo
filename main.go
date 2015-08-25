package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/codegangsta/cli"
	"github.com/msempere/remotgo/utils"
)

func main() {
	app := cli.NewApp()
	app.Name = "remotgo"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "role",
			Value: "role",
			Usage: "instance role",
		},
		cli.StringFlag{
			Name:  "environment",
			Value: "environment",
			Usage: "instance environment",
		},
		cli.StringFlag{
			Name:  "username",
			Value: "",
			Usage: "ssh username",
		},
		cli.StringFlag{
			Name:  "password",
			Value: "",
			Usage: "ssh password",
		},
		cli.StringFlag{
			Name:  "command",
			Value: "command",
			Usage: "command",
		},
	}
	app.Action = func(c *cli.Context) {
		ins, err := utils.GetInstances()

		if err != nil {
			panic(err)
		}

		username := c.String("username")

		if len(username) == 0 {
			user, err := user.Current()
			if err != nil {
				panic(err)
			}
			username = user.Username
		}

		instances := utils.Filter(ins, utils.CreateFilter(map[string]string{"role": c.String("role"), "environment": c.String("environment")}))
		for _, instance := range instances {
			fmt.Println(*instance.PublicDnsName)
			_, _, result, err := utils.SshExec(*instance.PublicDnsName, username, c.String("password"), c.String("command"), 100)
			if len(err) != 0 {
				fmt.Println(err)
			} else {
				fmt.Println(result)
			}
		}
	}
	app.Run(os.Args)

}

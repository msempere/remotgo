package main

import (
	"os"
	"os/user"
	"sync"

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
		cli.IntFlag{
			Name:  "timeout",
			Value: 100,
			Usage: "ssh command timeout",
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
		var wg sync.WaitGroup
		wg.Add(len(instances))

		for _, instance := range instances {
			go func() {
				defer wg.Done()
				_, _, result, err := utils.SshExec(*instance.PublicDnsName, username, c.String("password"), c.String("command"), c.Int("timeout"))

				if len(err) != 0 {
					utils.RenderOutput(*instance.PublicDnsName, err)
				} else {
					utils.RenderOutput(*instance.PublicDnsName, result)
				}
			}()
		}
		wg.Wait()
	}
	app.Run(os.Args)

}

package main

import (
	"os"
	"os/user"
	"sync"

	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/codegangsta/cli"
	"github.com/msempere/remotgo/utils"
)

func main() {
	app := cli.NewApp()
	app.Name = "remotgo"
	app.Email = "msempere@gmx.com"
	app.Usage = "Send commands over ssh to AWS EC2 instances"
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "role",
			Value: "role",
			Usage: "Instance role",
		},
		cli.StringFlag{
			Name:  "environment",
			Value: "environment",
			Usage: "Instance environment",
		},
		cli.StringFlag{
			Name: "username",
			Value: func() string {
				username, err := user.Current()
				if err != nil {
					return ""
				}
				return username.Username
			}(),
			Usage: "Ssh username (default: current user)",
		},
		cli.StringFlag{
			Name:  "password",
			Value: "",
			Usage: "Ssh password (default: empty)",
		},
		cli.StringFlag{
			Name:  "command",
			Value: "ls ~",
			Usage: "Command to execute.",
		},
		cli.BoolFlag{
			Name:  "quiet",
			Usage: "Quiet mode (default: false)",
		},
		cli.IntFlag{
			Name:  "timeout",
			Value: 200,
			Usage: "Shh command timeout (default: 200)",
		},
	}
	app.Action = func(c *cli.Context) {
		ins, err := utils.GetInstances()

		if err != nil {
			panic(err)
		}

		instances := utils.Filter(ins, utils.CreateFilter(map[string]string{"role": c.String("role"), "environment": c.String("environment")}))
		var wg sync.WaitGroup
		wg.Add(len(instances))

		for _, instance := range instances {
			go func(instance ec2.Instance) {
				defer wg.Done()
				_, _, result, err := utils.SshExec(*instance.PublicDnsName, c.String("username"), c.String("password"), c.String("command"), c.Int("timeout"))

				if len(err) != 0 {
					utils.RenderOutput(*instance.PublicDnsName, err, c.Bool("quiet"))
				} else {
					utils.RenderOutput(*instance.PublicDnsName, result, c.Bool("quiet"))
				}
			}(*instance)
		}
		wg.Wait()
	}
	app.Run(os.Args)

}

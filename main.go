package main

import (
	"os"
	"runtime"
	"sync"

	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/codegangsta/cli"
	"github.com/msempere/remotgo/utils"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	app := cli.NewApp()
	app.Name = "remotgo"
	app.Email = "msempere@gmx.com"
	app.Usage = "Send commands over ssh to AWS EC2 instances"
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "username, u",
			Value: utils.GetUsername(),
			Usage: "Ssh username (default: current user)",
		},
		cli.StringFlag{
			Name:  "password, p",
			Value: "",
			Usage: "Ssh password (default: empty)",
		},
		cli.StringFlag{
			Name:  "command, c",
			Value: "uname -a",
			Usage: "Command to execute.",
		},
		cli.BoolFlag{
			Name:  "quiet, q",
			Usage: "Quiet mode (default: false)",
		},
		cli.IntFlag{
			Name:  "timeout, o",
			Value: 200,
			Usage: "Shh command timeout (default: 200)",
		},
		cli.StringSliceFlag{
			Name:  "tags, t",
			Usage: "EC2 instance tags",
		},
		cli.StringFlag{
			Name:  "rsa, r",
			Value: utils.GetDefaultRSAFilePath(),
			Usage: "Path to RSA file (default ~/.ssh/id_rsa)",
		},
		cli.StringFlag{
			Name:  "dsa, d",
			Value: utils.GetDefaultDSAFilePath(),
			Usage: "Path to DSA file (default ~/.ssh/id_dsa)",
		},
	}
	app.Action = func(c *cli.Context) {
		instances, err := utils.GetInstances(c.StringSlice("tags"))

		if err != nil {
			panic(err)
		}
		var wg sync.WaitGroup
		wg.Add(len(instances))

		for _, instance := range instances {
			go func(instance ec2.Instance) {
				defer wg.Done()
				_, _, result, err := utils.SshExec(
					*instance.PublicDnsName,
					c.String("username"),
					c.String("password"),
					c.String("rsa"),
					c.String("dsa"),
					c.String("command"),
					c.Int("timeout"))

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

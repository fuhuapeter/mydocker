package container

import (
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/weikeit/mydocker/pkg/container"
)

var Init = cli.Command{
	Name:   "init",
	Usage:  "Run user's process in container. Do not call it outside!",
	Hidden: true,
	Action: func(ctx *cli.Context) error {
		log.Debugf("auto-calling initCommand...")
		return container.RunContainerInitProcess()
	},
}

var Run = cli.Command{
	Name:  "run",
	Usage: "Create a new mydocker container",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "detach,d",
			Usage: "Run the container in background",
		},
		cli.StringFlag{
			Name:  "name",
			Usage: "Assign a name to the container",
		},
		cli.StringFlag{
			Name:  "image",
			Usage: "The image to be used (alpine/busybox)",
		},
		cli.StringFlag{
			Name:  "memory,m",
			Usage: "Limit the container memory usage",
		},
		cli.Int64Flag{
			Name:  "cpu-period",
			Value: 250000,
			Usage: "Limit CPU CFS period in us",
		},
		cli.Int64Flag{
			Name:  "cpu-quota",
			Usage: "Limit CPU CFS quota in us",
		},
		cli.Int64Flag{
			Name:  "cpu-share",
			Usage: "CPU shares (relative weight)",
		},
		cli.StringFlag{
			Name:  "cpuset",
			Usage: "CPUs in which to allow execution (0-3, 0,1)",
		},
		cli.StringSliceFlag{
			Name:  "volume, v",
			Usage: "Bind a local directory/file, e.g. -v /src:/dst",
		},
		cli.StringSliceFlag{
			Name:  "env, e",
			Usage: "Set environment variables, e.g. -e key=value",
		},
		cli.StringFlag{
			Name:  "network, n",
			Usage: "Connect the container to a network",
		},
		cli.StringSliceFlag{
			Name:  "publish, p",
			Usage: "Publish the container's port(s) to the host",
		},
	},
	Action: func(ctx *cli.Context) error {
		c, err := container.NewContainer(ctx)
		if err != nil {
			return err
		}
		return c.Run()
	},
}

var List = cli.Command{
	Name:  "ps",
	Usage: "List all containers on the host",
	Action: func(ctx *cli.Context) error {
		return listContainers(ctx)
	},
}

var Logs = cli.Command{
	Name:  "logs",
	Usage: "Show all the logs of a container",
	Action: func(ctx *cli.Context) error {
		c, err := getContainerFromArg(ctx)
		if err != nil {
			return err
		}
		return c.Logs()
	},
}

var Exec = cli.Command{
	Name:  "exec",
	Usage: "Run a command in a running container",
	Action: func(ctx *cli.Context) error {
		c, cmdArray, err := parseExecArgs(ctx)
		if err != nil {
			return err
		}
		if c == nil {
			return nil
		}
		return c.Exec(cmdArray)
	},
}

var Stop = cli.Command{
	Name:  container.Stop,
	Usage: "Stop one or more containers",
	Action: func(ctx *cli.Context) error {
		return operateContainers(ctx, container.Stop)
	},
}

var Start = cli.Command{
	Name:  container.Start,
	Usage: "Start one or more containers",
	Action: func(ctx *cli.Context) error {
		return operateContainers(ctx, container.Start)
	},
}

var Restart = cli.Command{
	Name:  container.Restart,
	Usage: "Restart one or more containers",
	Action: func(ctx *cli.Context) error {
		return operateContainers(ctx, container.Restart)
	},
}

var Delete = cli.Command{
	Name:  "rm",
	Usage: "Delete one or more containers",
	Action: func(ctx *cli.Context) error {
		return operateContainers(ctx, container.Delete)
	},
}
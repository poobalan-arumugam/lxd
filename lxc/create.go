package main

import (
	"fmt"

	"github.com/lxc/lxd"
)

type createCmd struct{}

const createUsage = `
lxc create images:ubuntu <name>

Creates a container using the specified image and name
`

func (c *createCmd) usage() string {
	return createUsage
}

func (c *createCmd) flags() {}

func (c *createCmd) run(config *lxd.Config, args []string) error {
	if len(args) > 2 {
		return errArgs
	}

	if len(args) < 1 {
		return errArgs
	}

	if args[0] != "images:ubuntu" {
		return fmt.Errorf("Only the default ubuntu image is supported. Try `lxc create images:ubuntu foo`.")
	}

	var name string
	var remote string
	if len(args) == 2 {
		remote, name = config.ParseRemoteAndContainer(args[1])
	} else {
		name = ""
		remote = ""
	}

	d, err := lxd.NewClient(config, remote)
	if err != nil {
		return err
	}

	// TODO: implement the syntax for supporting other image types/remotes
	resp, err := d.Create(name)
	if err != nil {
		return err
	}

	return d.WaitForSuccess(resp.Operation)
}

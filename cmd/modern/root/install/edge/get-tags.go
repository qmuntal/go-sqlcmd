// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package edge

import (
	"github.com/microsoft/go-sqlcmd/internal/cmdparser"
	"github.com/microsoft/go-sqlcmd/internal/container"
	"github.com/microsoft/go-sqlcmd/internal/localizer"
)

type GetTags struct {
	cmdparser.Cmd
}

func (c *GetTags) DefineCommand(...cmdparser.CommandOptions) {
	options := cmdparser.CommandOptions{
		Use:   "get-tags",
		Short: localizer.Sprintf("Get tags available for Azure SQL Edge install"),
		Examples: []cmdparser.ExampleOptions{
			{
				Description: localizer.Sprintf("List tags"),
				Steps:       []string{"sqlcmd create azsql-edge get-tags"},
			},
		},
		Aliases: []string{"gt", "lt"},
		Run:     c.run,
	}

	c.Cmd.DefineCommand(options)
}

func (c *GetTags) run() {
	output := c.Output()

	tags := container.ListTags(
		"azure-sql-edge",
		"https://mcr.microsoft.com",
	)
	output.Struct(tags)
}

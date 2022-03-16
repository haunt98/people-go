package cli

import (
	"github.com/haunt98/people-go/internal/people"
	"github.com/urfave/cli/v2"
)

type action struct {
	peopleHandler people.Handler
}

func (a *action) RunHelp(c *cli.Context) error {
	return cli.ShowAppHelp(c)
}

func (a *action) RunList(c *cli.Context) error {
	return a.peopleHandler.List(c.Context)
}

func (a *action) RunAdd(c *cli.Context) error {
	return a.peopleHandler.Add(c.Context)
}

func (a *action) RunUpdate(c *cli.Context) error {
	return a.peopleHandler.Update(c.Context)
}

func (a *action) RunRemove(c *cli.Context) error {
	return a.peopleHandler.Remove(c.Context)
}

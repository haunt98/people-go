package cli

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/haunt98/people-go/internal/people"
	"github.com/make-go-great/color-go"
	"github.com/urfave/cli/v2"
)

const (
	Name  = "people"
	usage = "tracking people"

	commandList   = "list"
	commandGet    = "get"
	commandAdd    = "add"
	commandUpdate = "update"
	commandRemove = "remove"

	usageList   = "list people"
	usageGet    = "get person"
	usageAdd    = "add person"
	usageUpdate = "update person"
	usageRemove = "remove person"
)

type App struct {
	cliApp *cli.App
}

func NewApp(db *sql.DB, shouldInitDatabase bool, location *time.Location) (*App, error) {
	peopleRepo, err := people.NewRepository(context.Background(), db, shouldInitDatabase)
	if err != nil {
		return nil, fmt.Errorf("failed to new repository: %w", err)
	}
	peopleService := people.NewService(peopleRepo, location)
	peopleHandler := people.NewHandler(peopleService)

	a := &action{
		peopleHandler: peopleHandler,
	}

	cliApp := &cli.App{
		Name:   Name,
		Usage:  usage,
		Action: a.RunHelp,
		Commands: []*cli.Command{
			{
				Name:   commandList,
				Usage:  usageList,
				Action: a.RunList,
			},
			{
				Name:   commandAdd,
				Usage:  usageAdd,
				Action: a.RunAdd,
			},
			{
				Name:   commandUpdate,
				Usage:  usageUpdate,
				Action: a.RunUpdate,
			},
			{
				Name:   commandRemove,
				Usage:  usageRemove,
				Action: a.RunRemove,
			},
		},
	}

	return &App{
		cliApp: cliApp,
	}, nil
}

func (a *App) Run() {
	if err := a.cliApp.Run(os.Args); err != nil {
		color.PrintAppError(Name, err.Error())
	}
}

package cli

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/urfave/cli/v2"

	"github.com/make-go-great/color-go"

	"github.com/haunt98/people-go/internal/people"
)

const (
	Name  = "people"
	usage = "tracking people"

	commandList   = "list"
	commandAdd    = "add"
	commandUpdate = "update"
	commandRemove = "remove"
	commandExport = "export"
	commandImport = "import"
	commandDrop   = "drop"

	usageList   = "list people"
	usageAdd    = "add person"
	usageUpdate = "update person"
	usageRemove = "remove person"
	usageExport = "export data"
	usageImport = "import data"
	usageDrop   = "drop all data (be careful)"
)

type App struct {
	cliApp *cli.App
}

func NewApp(db *sql.DB, location *time.Location) (*App, error) {
	peopleRepo, err := people.NewRepository(context.Background(), db)
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
			{
				Name:   commandExport,
				Usage:  usageExport,
				Action: a.RunExport,
			},
			{
				Name:   commandImport,
				Usage:  usageImport,
				Action: a.RunImport,
			},
			{
				Name:   commandDrop,
				Usage:  usageDrop,
				Action: a.RunDrop,
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

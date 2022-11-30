package main

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"
	"time"

	_ "modernc.org/sqlite"

	"github.com/make-go-great/color-go"
	"github.com/make-go-great/xdg-go"

	"github.com/haunt98/people-go/internal/cli"
)

const (
	appName      = "people"
	dataFilename = "data.sqlite3"
)

func main() {
	if err := os.MkdirAll(getDataDirPath(), 0o755); err != nil {
		log.Fatalln(err)
	}

	db, err := sql.Open("sqlite", getDataFilePath())
	if err != nil {
		color.PrintAppError(appName, err.Error())
		return
	}
	defer db.Close()

	// Shout out to Sai Gon, Viet Nam
	location, err := time.LoadLocation("Asia/Ho_Chi_Minh")
	if err != nil {
		color.PrintAppError(appName, err.Error())
		return
	}

	app, err := cli.NewApp(db, location)
	if err != nil {
		color.PrintAppError(appName, err.Error())
		return
	}

	app.Run()
}

// Should be ~/.local/share/people
func getDataDirPath() string {
	return filepath.Join(xdg.GetDataHome(), cli.Name)
}

// Should be ~/.local/share/people/data.sqlite3
func getDataFilePath() string {
	return filepath.Join(xdg.GetDataHome(), cli.Name, dataFilename)
}

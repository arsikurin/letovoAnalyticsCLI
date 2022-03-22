package main

import (
	"github.com/arsikurin/letovoAnalyticsCLI/src/cmd"
	"github.com/arsikurin/letovoAnalyticsCLI/src/utils/database"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(5) // 1 for prod
	err := database.CreateDatabaseIfNotExists("sqlite3", "./db.sqlite")
	if err != nil {
		log.Fatalln(err)
	}

	err = cmd.Execute()
	if err != nil {
		log.Errorln(err)
	}
}

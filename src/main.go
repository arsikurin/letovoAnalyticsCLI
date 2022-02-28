package main

import (
	"database/sql"
	"github.com/arsikurin/letovoAnalyticsCLI/src/cmd"
	"github.com/arsikurin/letovoAnalyticsCLI/src/utils/database"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
)

func main() {
	db, err := database.OpenFileDB("sqlite3", "./db.sqlite")
	if err != nil {
		log.Fatalln(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Errorln(err)
		}
	}(db)

	err = cmd.Execute()
	if err != nil {
		log.Errorln(err)
	}
}

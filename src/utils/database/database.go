package database

import (
	"database/sql"
	"errors"
	"github.com/arsikurin/letovoAnalyticsCLI/src/utils/types"
	log "github.com/sirupsen/logrus"
	"os"
)

func CreateDatabaseIfNotExists(driverName string, dataSourceName string) error {
	if _, err := os.Stat(dataSourceName); errors.Is(err, os.ErrNotExist) {
		file, err := os.Create(dataSourceName)
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				log.Errorln(err)
			}
		}(file)
		if err != nil {
			return err
		}
		log.Debug(dataSourceName, " created")

		db, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return err
		}
		err = createTable(db)
		if err != nil {
			return err
		}
		err = initUser(db)
		if err != nil {
			return err
		}
		log.Debug(dataSourceName, " inited")
		defer func(db *sql.DB) {
			err := db.Close()
			if err != nil {
				log.Errorln(err)
			}
		}(db)
	}
	return nil
}

func OpenFileDB(driverName string, dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}

	if il, err := isLogged(db); !il {
		if err != nil {
			return nil, err
		}
		return nil, errors.New("[✘] Credentials not found in the database. Consider registering")
	}
	return db, nil
}

func isLogged(db *sql.DB) (bool, error) {
	row, err := db.Query(
		"SELECT password, login FROM users",
	)
	if err != nil {
		return false, err
	}
	defer func(row *sql.Rows) {
		err := row.Close()
		if err != nil {
			log.Errorln(err)
		}
	}(row)

	for row.Next() {
		var (
			password *string
			login    *string
		)
		err := row.Scan(&password, &login)
		if err != nil {
			return false, err
		}
		if password != nil && login != nil {
			return true, nil
		} else {
			return false, nil
		}
	}
	return false, errors.New("nothing found in the database")
}

func createTable(db *sql.DB) error {
	query, err := db.Prepare(
		`CREATE TABLE users (
		id INTEGER PRIMARY KEY,		
		student_id INTEGER,		
		token VARCHAR(255),
		password VARCHAR(255),
		login VARCHAR(255)	
	  );`,
	)
	if err != nil {
		return err
	}

	_, err = query.Exec()
	if err != nil {
		return err
	}
	return nil
}

func initUser(db *sql.DB) error {
	query, err := db.Prepare(
		`INSERT INTO users (student_id, token, password, login) VALUES (?, ?, ?, ?)`,
	)
	if err != nil {
		return err
	}

	_, err = query.Exec(nil, nil, nil, nil)
	if err != nil {
		return err
	}
	return nil
}

func UpdateCreds(db *sql.DB, password string, login string, studentID int) error {
	//goland:noinspection SqlWithoutWhere
	query, err := db.Prepare(
		`UPDATE users SET password=?, login=?, student_id=?`,
	)
	if err != nil {
		return err
	}

	_, err = query.Exec(password, login, studentID)
	if err != nil {
		return err
	}
	return nil
}

func GetCreds(db *sql.DB) (types.DatabaseCreds, error) {
	row, err := db.Query(
		"SELECT student_id, password, login FROM users",
	)
	if err != nil {
		return types.DatabaseCreds{}, err
	}
	defer func(row *sql.Rows) {
		err := row.Close()
		if err != nil {
			log.Errorln(err)
		}
	}(row)

	for row.Next() {
		var (
			studentID *int
			password  *string
			login     *string
		)
		err := row.Scan(&studentID, &password, &login)
		if err != nil {
			return types.DatabaseCreds{}, err
		}

		return types.DatabaseCreds{
			StudentID: studentID,
			Password:  password,
			Login:     login,
		}, nil
	}
	return types.DatabaseCreds{}, errors.New("nothing found in the database")
}

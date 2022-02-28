package database

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/arsikurin/letovoAnalyticsCLI/src/utils/types"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh/terminal"
	"os"
)

func OpenFileDB(driverName string, dataSourceName string) (*sql.DB, error) {
	var (
		err      error
		dbExists = true
	)

	if _, err := os.Stat(dataSourceName); errors.Is(err, os.ErrNotExist) {
		log.Println("Creating", dataSourceName)
		file, err := os.Create(dataSourceName)
		if err != nil {
			return nil, err
		}
		err = file.Close()
		if err != nil {
			log.Errorln(err)
		}
		log.Println(dataSourceName, "created")
		dbExists = false
	}

	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}

	if !dbExists {
		err = createTable(db)
		if err != nil {
			return nil, err
		}
		err = initUser(db)
		if err != nil {
			return nil, err
		}
	} else {
		il, err := isLogged(db)
		if err != nil {
			return nil, err
		}
		if il == false {
			//var (
			//	login    string
			//	password string
			//)
			fmt.Print("->")
			res, err := terminal.ReadPassword(0)
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println(string(res))
			//fmt.Println(getPassword("pass->"))
			//fmt.Println(getPassword("login->"))
		}
	}
	return db, nil
}

//// getPassword - Prompt for password. Use stty to disable echoing.
//func getPassword(prompt string) string {
//	fmt.Print(prompt)
//
//	// Common settings and variables for both stty calls.
//	attrs := syscall.ProcAttr{
//		Dir:   "",
//		Env:   []string{},
//		Files: []uintptr{os.Stdin.Fd(), os.Stdout.Fd(), os.Stderr.Fd()},
//		Sys:   nil}
//	var ws syscall.WaitStatus
//
//	// Disable echoing.
//	pid, err := syscall.ForkExec(
//		"/bin/stty",
//		[]string{"stty", "-echo"},
//		&attrs)
//	if err != nil {
//		panic(err)
//	}
//
//	// Wait for the stty process to complete.
//	_, err = syscall.Wait4(pid, &ws, 0, nil)
//	if err != nil {
//		panic(err)
//	}
//
//	// Echo is disabled, now grab the data.
//	reader := bufio.NewReader(os.Stdin)
//	text, err := reader.ReadString('\n')
//	if err != nil {
//		panic(err)
//	}
//
//	// Re-enable echo.
//	pid, err = syscall.ForkExec(
//		"/bin/stty",
//		[]string{"stty", "echo"},
//		&attrs)
//	if err != nil {
//		panic(err)
//	}
//
//	// Wait for the stty process to complete.
//	_, err = syscall.Wait4(pid, &ws, 0, nil)
//	if err != nil {
//		panic(err)
//	}
//
//	return strings.TrimSpace(text)
//}

func isLogged(db *sql.DB) (bool, error) {
	row, err := db.Query(
		"SELECT student_id, password, login FROM users",
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
			studentID *int
			password  *string
			login     *string
		)
		err := row.Scan(&studentID, &password, &login)
		if err != nil {
			return false, err
		}
		if studentID != nil && password != nil && login != nil {
			return true, nil
		} else {
			return false, nil
		}
	}
	return false, errors.New("nothing found in the database")
}

func createTable(db *sql.DB) error {
	stmt, err := db.Prepare(
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

	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	return nil
}

func initUser(db *sql.DB) error {
	statement, err := db.Prepare(
		`INSERT INTO users (student_id, token, password, login) VALUES (?, ?, ?, ?)`,
	)
	if err != nil {
		return err
	}

	_, err = statement.Exec(nil, nil, nil, nil)
	if err != nil {
		return err
	}
	return nil
}

func UpdateCreds(db *sql.DB, password string, login string) error {
	statement, err := db.Prepare(
		`UPDATE users SET password=?, login=?`,
	)
	if err != nil {
		return err
	}

	_, err = statement.Exec(password, login)
	if err != nil {
		return err
	}
	return nil
}

func GetData(db *sql.DB) (types.DatabaseCreds, error) {
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

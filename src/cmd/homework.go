package cmd

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/arsikurin/letovoAnalyticsCLI/src/utils/database"
	"html"
	"os"
	"time"

	"github.com/arsikurin/letovoAnalyticsCLI/src/utils"
	"github.com/arsikurin/letovoAnalyticsCLI/src/utils/api"
	"github.com/arsikurin/letovoAnalyticsCLI/src/utils/colorlib"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(homeworkCmd)

	homeworkCmd.Flags().BoolP("week", "w", false, "send homework for the week")
	homeworkCmd.Flags().StringP("day", "d", "", "specify a day")
}

var homeworkCmd = &cobra.Command{
	Use:     "homework",
	Aliases: []string{"hw", "h"},
	Short:   "Get homework (default is for today)",
	Long:    `Get homework from s.letovo.ru (default is for today)`,
	RunE: func(cmd *cobra.Command, args []string) error {
		entireStatus, err := cmd.Flags().GetBool("week")
		if err != nil {
			return err
		}
		specificDay, err := cmd.Flags().GetString("day")
		if err != nil {
			return err
		}
		db, err := database.OpenFileDB("sqlite3", "./db.sqlite")
		if err != nil {
			return err
		}
		defer func(db *sql.DB) {
			err := db.Close()
			if err != nil {
				log.Errorln(err)
			}
		}(db)

		if entireStatus {
			return sendWeekHomework()
		} else if specificDay != "" {
			day, err := utils.ParseDay(specificDay)
			if err != nil {
				return err
			}
			return sendHomeworkFor(db, day)
		} else {
			return sendHomeworkFor(db, time.Now().Weekday())
		}
	},
}

func sendWeekHomework() error {
	return errors.New("not implemented")
	//api.ReceiveScheduleAndHw(true, time.Monday)
}

func sendHomeworkFor(db *sql.DB, specificDay time.Weekday) error {
	if specificDay == time.Sunday {
		fmt.Println("Sunday! No lessons")
		os.Exit(0)
	}
	homeworkResponse, err := api.ReceiveScheduleAndHw(db, false, specificDay)
	if err != nil {
		return err
	}
	startOfWeek, err := time.Parse("2006-01-02", homeworkResponse.Data[0].Date)
	if err != nil {
		return errors.New("cannot parse time from response. Perhaps the layout has been changed")
	}
	fmt.Printf("%s%s, %s%s\n", colorlib.Style.Italic, specificDay, startOfWeek.Format("02.01.2006"), colorlib.Style.Reset)

	var (
		subject string
		flag    bool
	)
	for _, day := range homeworkResponse.Data {
		if len(day.Schedules) > 0 {
			if day.Schedules[0].Group.Subject.SubjectNameEng != "" {
				subject = day.Schedules[0].Group.Subject.SubjectNameEng
			} else {
				subject = day.Schedules[0].Group.Subject.SubjectName
			}

			flag = false
			payload := "\n" + day.PeriodName + ": " + colorlib.Fg.Red + colorlib.Style.Bold + subject + colorlib.Style.Reset + "\n"
			if len(day.Schedules[0].Lessons) > 0 {
				if day.Schedules[0].Lessons[0].LessonHw != "" {
					payload += day.Schedules[0].Lessons[0].LessonHw + "\n"
				} else {
					payload += colorlib.Style.Italic + "No homework\n" + colorlib.Style.Reset
				}

				if day.Schedules[0].Lessons[0].LessonUrl != "" {
					flag = true
					payload += colorlib.Fg.Blue + day.Schedules[0].Lessons[0].LessonUrl + colorlib.Fg.Reset + "\n"
				}
				if day.Schedules[0].Lessons[0].LessonHwUrl != "" {
					flag = true
					payload += colorlib.Fg.Blue + day.Schedules[0].Lessons[0].LessonHwUrl + colorlib.Fg.Reset + "\n"
				}
				if !flag {
					payload += colorlib.Style.Italic + "No links attached\n" + colorlib.Style.Reset
				}
				if day.Schedules[0].Lessons[0].LessonTopic != "" {
					payload += day.Schedules[0].Lessons[0].LessonTopic
				} else {
					payload += colorlib.Style.Italic + "No topic" + colorlib.Style.Reset
				}
			} else {
				payload += "Lessons not found"
			}
			fmt.Println(html.UnescapeString(payload))
		}
	}
	return nil
}

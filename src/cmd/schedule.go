package cmd

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/arsikurin/letovoAnalyticsCLI/src/utils"
	"github.com/arsikurin/letovoAnalyticsCLI/src/utils/api"
	"github.com/arsikurin/letovoAnalyticsCLI/src/utils/colorlib"
	"github.com/arsikurin/letovoAnalyticsCLI/src/utils/database"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"html"
	"os"
	"time"
)

func init() {
	rootCmd.AddCommand(scheduleCmd)

	scheduleCmd.Flags().BoolP("week", "w", false, "send schedule for the week")
	scheduleCmd.Flags().StringP("day", "d", "", "specify a day")
}

var scheduleCmd = &cobra.Command{
	Use:     "schedule",
	Aliases: []string{"sch", "s"},
	Short:   "Get schedule (default is for today)",
	Long:    `Get schedule from s.letovo.ru (default is for today)`,
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
			return sendWeekSchedule()
		} else if specificDay != "" {
			day, err := utils.ParseDay(specificDay)
			if err != nil {
				return err
			}
			return sendScheduleFor(db, day)
		} else {
			return sendScheduleFor(db, time.Now().Weekday())
		}
	},
}

func sendWeekSchedule() error {
	return errors.New("not implemented")
	//api.ReceiveScheduleAndHw(true, time.Monday)
}

func sendScheduleFor(db *sql.DB, specificDay time.Weekday) error {
	if specificDay == time.Sunday {
		fmt.Println("Sunday! No lessons")
		os.Exit(0)
	}
	scheduleResponse, err := api.ReceiveScheduleAndHw(db, false, specificDay)
	if err != nil {
		return err
	}
	startOfWeek, err := time.Parse("2006-01-02", scheduleResponse.Data[0].Date)
	if err != nil {
		return errors.New("cannot parse time from response. Perhaps the layout has been changed")
	}
	fmt.Printf("%s%s, %s%s\n", colorlib.Style.Italic, specificDay, startOfWeek.Format("02.01.2006"), colorlib.Style.Reset)

	var subject string
	for _, day := range scheduleResponse.Data {
		if len(day.Schedules) > 0 {
			payload := "\n" + day.PeriodName + " | " + colorlib.Style.Italic + day.Schedules[0].Room.RoomName + colorlib.Style.Reset + ":\n"
			if day.Schedules[0].Group.Subject.SubjectNameEng != "" {
				subject = day.Schedules[0].Group.Subject.SubjectNameEng
			} else {
				subject = day.Schedules[0].Group.Subject.SubjectName
			}
			payload += colorlib.Fg.Red + colorlib.Style.Bold + subject + colorlib.Fg.Reset + " " + day.Schedules[0].Group.GroupName + "\n"
			payload += day.PeriodStart + " — " + day.PeriodEnd

			if day.Schedules[0].ZoomMeetings != nil {
				payload += "\n" + colorlib.Fg.Blue + day.Schedules[0].ZoomMeetings.MeetingUrl + colorlib.Fg.Reset
			}
			fmt.Println(html.UnescapeString(payload))
		}
	}
	return nil
}

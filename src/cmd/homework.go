package cmd

import (
	"fmt"
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
	Run: func(cmd *cobra.Command, args []string) {
		// get the flag value, its default value is false
		entireStatus, _ := cmd.Flags().GetBool("week")
		specificDay, _ := cmd.Flags().GetString("day")

		if entireStatus {
			sendWeekHomework()
		} else if specificDay != "" {
			sendHomeworkFor(utils.ParseDay(specificDay))
		} else {
			sendHomeworkFor(time.Now().Weekday())
		}
	},
}

func sendWeekHomework() {
	log.Fatalln("Not implemented!")
	api.ReceiveScheduleAndHw(true)
}

func sendHomeworkFor(specificDay time.Weekday) {
	if specificDay == time.Sunday {
		fmt.Println("Sunday! No lessons")
		os.Exit(0)
	}
	homeworkResponse := api.ReceiveScheduleAndHw(false)
	startOfWeek, err := time.Parse("2006-01-02", homeworkResponse.Data[0].Date)
	if err != nil {
		log.Errorln("cannot parse time from response. Perhaps the layout has been changed?")
	}
	fmt.Printf("%s%s %s%s\n", colorlib.Style.Italic, specificDay, startOfWeek.Format("02.01.2006"), colorlib.Style.Reset)

	for _, day := range homeworkResponse.Data {
		if len(day.Schedules) > 0 {
			// TODO
			payload := "\n" + day.PeriodName + " | " + colorlib.Style.Italic + day.Schedules[0].Room.RoomName + colorlib.Style.Reset + ":\n"
			var subject string
			if day.Schedules[0].Group.Subject.SubjectNameEng != nil {
				subject = *day.Schedules[0].Group.Subject.SubjectNameEng
			} else {
				subject = day.Schedules[0].Group.Subject.SubjectName
			}
			payload += colorlib.Fg.Red + subject + colorlib.Fg.Reset + " " + day.Schedules[0].Group.GroupName + "\n"
			payload += day.PeriodStart + " — " + day.PeriodEnd + "\n"

			if day.Schedules[0].ZoomMeetings != nil {
				payload += colorlib.Fg.Blue + day.Schedules[0].ZoomMeetings.MeetingUrl + colorlib.Fg.Reset
			}
			fmt.Println(payload)
		}
	}
}

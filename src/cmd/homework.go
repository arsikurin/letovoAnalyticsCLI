package cmd

import (
	"fmt"
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
	api.ReceiveScheduleAndHw(true, time.Monday)
}

func sendHomeworkFor(specificDay time.Weekday) {
	if specificDay == time.Sunday {
		fmt.Println("Sunday! No lessons")
		os.Exit(0)
	}
	homeworkResponse := api.ReceiveScheduleAndHw(false, specificDay)
	startOfWeek, err := time.Parse("2006-01-02", homeworkResponse.Data[0].Date)
	if err != nil {
		log.Errorln("Cannot parse time from response. Perhaps the layout has been changed?")
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
}

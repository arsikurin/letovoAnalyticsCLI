package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/arsikurin/letovoAnalyticsCLI/src/utils"
	"github.com/arsikurin/letovoAnalyticsCLI/src/utils/colorlib"
	"github.com/arsikurin/letovoAnalyticsCLI/src/utils/types"

	"github.com/spf13/cobra"
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
	Run: func(cmd *cobra.Command, args []string) {
		// get the flag value, its default value is false
		entireStatus, _ := cmd.Flags().GetBool("week")
		specificDay, _ := cmd.Flags().GetString("day")

		if entireStatus {
			sendWeekSchedule()
		} else if specificDay != "" {
			sendScheduleFor(parseDay(specificDay))
		} else {
			sendScheduleFor(time.Now().Weekday())
		}
	},
}

func parseDay(dayRaw string) time.Weekday {
	isMonday, _ := regexp.MatchString(`(?i)^mo`, dayRaw)
	isTuesday, _ := regexp.MatchString(`(?i)^tu`, dayRaw)
	isWednesday, _ := regexp.MatchString(`(?i)^we`, dayRaw)
	isThursday, _ := regexp.MatchString(`(?i)^th`, dayRaw)
	isFriday, _ := regexp.MatchString(`(?i)^fr`, dayRaw)
	isSaturday, _ := regexp.MatchString(`(?i)^sa`, dayRaw)

	switch true {
	case isMonday:
		return time.Monday
	case isTuesday:
		return time.Tuesday
	case isWednesday:
		return time.Wednesday
	case isThursday:
		return time.Thursday
	case isFriday:
		return time.Friday
	case isSaturday:
		return time.Saturday
	default:
		utils.Err("Cannot parse specific day")
		return -1
	}
}

func sendWeekSchedule() {
	utils.Err("Not implemented!")
	receiveScheduleAndHw(true)
}

func sendScheduleFor(specificDay time.Weekday) {
	if specificDay == time.Sunday {
		fmt.Println("Sunday! No lessons")
		os.Exit(0)
	}
	scheduleResponse := receiveScheduleAndHw(false)
	startOfWeek, _ := time.Parse("2006-01-02", scheduleResponse.Data[0].Date)
	fmt.Printf("%s%s %s%s\n", colorlib.Style.Italic, specificDay, startOfWeek.Format("02.01.2006"), colorlib.Style.Reset)

	for _, day := range scheduleResponse.Data {
		if len(day.Schedules) > 0 {
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

func receiveScheduleAndHw(week bool) types.ScheduleResponse {
	var (
		urlAddr   string
		date      string
		studentId = 54405
	)

	if time.Now().Weekday() == 0 {
		date = time.Now().Add(time.Hour * 24).Format("2006-01-02")
	} else {
		date = time.Now().Format("2006-01-02")
	}

	if week {
		urlAddr = "https://s-api.letovo.ru/api/schedule/" + strconv.Itoa(studentId) + "/week?schedule_date=" + date
	} else {
		urlAddr = "https://s-api.letovo.ru/api/schedule/" + strconv.Itoa(studentId) + "/day?schedule_date=" + date
	}

	client := &http.Client{}
	req, _ := http.NewRequest("GET", urlAddr, nil)
	req.Header.Set("Authorization", "Bearer "+receiveToken().Data.Token)
	resp, err := client.Do(req)
	if err != nil {
		utils.Err(err)
	}
	if resp.StatusCode != 200 {
		utils.Err(resp.Status)
	}
	//var v interface{}
	//body, err := ioutil.ReadAll(resp.Body)
	//json.Unmarshal(body, &v)
	//fmt.Println(v)
	res := new(types.ScheduleResponse)
	err = json.NewDecoder(resp.Body).Decode(res)
	if err != nil {
		utils.Err(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			utils.Err(err)
		}
	}(resp.Body)
	//jsonData, err := json.MarshalIndent(res, "", "    ")
	//fmt.Println(string(jsonData))
	return *res
}

func receiveToken() types.TokenResponse {
	data := url.Values{
		"login":    {"login"},
		"password": {"pass"},
	}

	resp, err := http.PostForm("https://s-api.letovo.ru/api/login", data)
	if err != nil {
		utils.Err(err)
	}
	if resp.StatusCode != 200 {
		utils.Err("unable to obtain token: " + resp.Status)
	}

	res := new(types.TokenResponse)
	err = json.NewDecoder(resp.Body).Decode(res)
	if err != nil {
		utils.Err(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			utils.Err(err)
		}
	}(resp.Body)
	return *res
}

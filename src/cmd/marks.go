package cmd

import (
	"database/sql"
	"fmt"
	"github.com/arsikurin/letovoAnalyticsCLI/src/utils/api"
	"github.com/arsikurin/letovoAnalyticsCLI/src/utils/colorlib"
	"github.com/arsikurin/letovoAnalyticsCLI/src/utils/database"
	"github.com/arsikurin/letovoAnalyticsCLI/src/utils/types"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"html"
	"strconv"
	"time"
)

func init() {
	rootCmd.AddCommand(marksCmd)

	marksCmd.Flags().BoolP("all", "a", false, "send all marks")
	marksCmd.Flags().BoolP("final", "f", false, "send final marks")
	marksCmd.Flags().BoolP("summative", "s", false, "send summative marks")
}

var marksCmd = &cobra.Command{
	Use:     "marks",
	Aliases: []string{"ma", "m"},
	Short:   "Get marks (default is for today)",
	Long:    `Get marks from s.letovo.ru (default is marks within one week)`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// get the flag value, its default value is false
		allStatus, err := cmd.Flags().GetBool("all")
		if err != nil {
			return err
		}
		finalStatus, err := cmd.Flags().GetBool("final")
		if err != nil {
			return err
		}
		summativeStatus, err := cmd.Flags().GetBool("summative")
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

		if allStatus {
			return sendAllMarks(db)
		} else if finalStatus {
			return sendFinalMarks(db)
		} else if summativeStatus {
			return sendSummativeMarks(db)
		} else {
			return sendRecentMarks(db)
		}
	},
}

func sendAllMarks(db *sql.DB) error {
	marksResponse, err := api.ReceiveMarks(db)
	if err != nil {
		return err
	}
	for _, subject := range marksResponse.Data {
		payload := colorlib.Style.Bold + colorlib.Fg.Red + subject.Group.Subject.SubjectNameEng + colorlib.Style.Reset + "\n"
		if len(subject.FormativeList) > 0 {
			for _, mark := range subject.FormativeList {
				payload += colorlib.Style.Bold + string(mark.MarkValue) + colorlib.Style.Reset + "F "
			}
		}
		if len(subject.SummativeList) > 0 {
			err := prepareSummativeMarks(subject, &payload, false)
			if err != nil {
				return err
			}

		}
		if len(subject.FormativeList) > 0 || len(subject.SummativeList) > 0 {
			fmt.Println("\n" + html.UnescapeString(payload))
		}
	}
	return nil
}
func sendSummativeMarks(db *sql.DB) error {
	marksResponse, err := api.ReceiveMarks(db)
	if err != nil {
		return err
	}
	for _, subject := range marksResponse.Data {
		if len(subject.SummativeList) > 0 {
			payload := colorlib.Style.Bold + colorlib.Fg.Red + subject.Group.Subject.SubjectNameEng + colorlib.Style.Reset + "\n"
			err := prepareSummativeMarks(subject, &payload, false)
			if err != nil {
				return err
			}
			fmt.Println("\n" + html.UnescapeString(payload))
		}
	}
	return nil
}
func sendFinalMarks(db *sql.DB) error {
	marksResponse, err := api.ReceiveMarks(db)
	if err != nil {
		return err
	}
	for _, subject := range marksResponse.Data {
		if len(subject.FinalMarkList) > 0 {
			payload := colorlib.Style.Bold + colorlib.Fg.Red + subject.Group.Subject.SubjectNameEng + colorlib.Style.Reset + "\n"
			for _, mark := range subject.FinalMarkList {
				payload += colorlib.Style.Bold + mark.FinalValue + colorlib.Style.Reset + mark.FinalCriterion + " "
			}
			if subject.GroupAvgMark != "" {
				payload += " | " + colorlib.Style.Italic + "group_avg: " + colorlib.Style.Reset + colorlib.Style.Bold + subject.GroupAvgMark + colorlib.Style.Reset
			}
			if subject.ResultFinalMark != "" {
				payload += " | " + colorlib.Style.Italic + "final: " + colorlib.Style.Reset + colorlib.Style.Bold + subject.ResultFinalMark + colorlib.Style.Reset
			}
			fmt.Println("\n" + html.UnescapeString(payload))
		}
	}
	return nil
}
func sendRecentMarks(db *sql.DB) error {
	marksResponse, err := api.ReceiveMarks(db)
	if err != nil {
		return err
	}
	var flag bool
	for _, subject := range marksResponse.Data {
		flag = false
		payload := colorlib.Style.Bold + colorlib.Fg.Red + subject.Group.Subject.SubjectNameEng + colorlib.Style.Reset + "\n"
		if len(subject.FormativeList) > 0 {
			for _, mark := range subject.FormativeList {
				createdAt, err := time.Parse("2006-01-02 15:04:05", mark.CreatedAt)
				if err != nil {
					return err
				}
				if time.Now().Sub(createdAt).Hours()/24 <= 8 {
					flag = true
					payload += colorlib.Style.Bold + string(mark.MarkValue) + colorlib.Style.Reset + "F "
				}
			}
		}
		if len(subject.SummativeList) > 0 {
			err := prepareSummativeMarks(subject, &payload, true)
			if err != nil {
				return err
			}
		}
		if flag || len(subject.SummativeList) > 0 {
			fmt.Println("\n" + html.UnescapeString(payload))
		}
	}
	return nil
}

func prepareSummativeMarks(subject types.Subject, payload *string, checkDate bool) error {
	markA, markB, markC, markD := [2]int{0, 0}, [2]int{0, 0}, [2]int{0, 0}, [2]int{0, 0}
	for _, mark := range subject.SummativeList {
		if mark.MarkValue.IsDigit() {
			switch mark.MarkCriterion {
			case "A":
				mv, err := mark.MarkValue.Int()
				if err != nil {
					return err
				}
				markA[0] += mv
				markA[1] += 1
			case "B":
				mv, err := mark.MarkValue.Int()
				if err != nil {
					return err
				}
				markB[0] += mv
				markB[1] += 1
			case "C":
				mv, err := mark.MarkValue.Int()
				if err != nil {
					return err
				}
				markC[0] += mv
				markC[1] += 1
			case "D":
				mv, err := mark.MarkValue.Int()
				if err != nil {
					return err
				}
				markD[0] += mv
				markD[1] += 1
			}
		}
		createdAt, err := time.Parse("2006-01-02 15:04:05", mark.CreatedAt)
		if err != nil {
			return err
		}
		if !checkDate || time.Now().Sub(createdAt).Hours()/24 <= 8 {
			*payload += colorlib.Style.Bold + string(mark.MarkValue) + colorlib.Style.Reset + mark.MarkCriterion + " "
		}
	}
	if markA[1] == 0 {
		markA[1] = 1
	}
	if markB[1] == 0 {
		markB[1] = 1
	}
	if markC[1] == 0 {
		markC[1] = 1
	}
	if markD[1] == 0 {
		markD[1] = 1
	}
	markAAvg, markBAvg := markA[0]/markA[1], markB[0]/markB[1]
	markCAvg, markDAvg := markC[0]/markC[1], markD[0]/markD[1]

	temp := *payload
	if temp[len(temp)-2] == 109 {
		*payload += "no recent marks"
	}

	*payload += " | " + colorlib.Style.Italic + "avg: " + colorlib.Style.Reset
	if float32(markAAvg) > 0 {
		*payload += colorlib.Style.Bold + strconv.Itoa(markAAvg) + colorlib.Style.Reset + "A "
	}
	if float32(markBAvg) > 0 {
		*payload += colorlib.Style.Bold + strconv.Itoa(markBAvg) + colorlib.Style.Reset + "B "
	}
	if float32(markCAvg) > 0 {
		*payload += colorlib.Style.Bold + strconv.Itoa(markCAvg) + colorlib.Style.Reset + "C "
	}
	if float32(markDAvg) > 0 {
		*payload += colorlib.Style.Bold + strconv.Itoa(markDAvg) + colorlib.Style.Reset + "D "
	}
	return nil
}

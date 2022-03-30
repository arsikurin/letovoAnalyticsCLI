package types

import (
	"strconv"
)

type DatabaseCreds struct {
	StudentID *int
	Password  *string
	Login     *string
}

type TokenResponse struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ExpiresAt string `json:"expires_at"`
		Token     string `json:"token"`
		TokenType string `json:"token_type"`
	} `json:"data"`
}

type MeResponse struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		User struct {
			StudentID int `json:"student_id"`
		} `json:"user"`
	} `json:"data"`
}

type (
	ScheduleAndHwResponse struct {
		Status  string `json:"status"`
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    []data `json:"data"`
	}
	data struct {
		PeriodNumDay int        `json:"period_num_day"`
		PeriodName   string     `json:"period_name"`
		PeriodStart  string     `json:"period_start"`
		PeriodEnd    string     `json:"period_end"`
		Date         string     `json:"date"`
		Schedules    []schedule `json:"schedules"`
	}
	schedule struct {
		ScheduleRoom int      `json:"schedule_room"`
		Lessons      []lesson `json:"lessons"`
		ZoomMeetings *struct {
			MeetingUrl string `json:"meeting_url"`
		} `json:"zoom_meetings"`
		Group group `json:"group"`
		Room  struct {
			RoomName        string `json:"room_name"`
			RoomDescription string `json:"room_description"`
		} `json:"room"`
	}
	lesson struct {
		LessonTopic   string `json:"lesson_thema,omitempty"`
		LessonUrl     string `json:"lesson_url,omitempty"`
		LessonHw      string `json:"lesson_hw,omitempty"`
		LessonHwDate  string `json:"lesson_hw_date,omitempty"`
		LessonHwUrl   string `json:"lesson_hw_url,omitempty"`
		LessonComment string `json:"lesson_comment,omitempty"`
		Attendance    []struct {
			AttendanceReason *string `json:"attendance_reason,omitempty"`
		} `json:"attendance"`
	}
	group struct {
		GroupName string `json:"group_name"`
		Subject   struct {
			SubjectName    string `json:"subject_name"`
			SubjectNameEng string `json:"subject_name_eng,omitempty"`
		} `json:"subject"`
	}
)

// MarksResponse represents ...
type (
	MarksResponse struct {
		Status  string    `json:"status"`
		Code    int       `json:"code"`
		Message string    `json:"message"`
		Data    []Subject `json:"data"`
	}
	Subject struct {
		FormativeAvgValue float32     `json:"formative_avg_value,omitempty"`
		SummativeAvgValue float32     `json:"summative_avg_value,omitempty"`
		FormativeList     []mark      `json:"formative_list"`
		SummativeList     []mark      `json:"summative_list"`
		GroupAvgMark      string      `json:"group_avg_mark,omitempty"`
		FinalMarkList     []finalMark `json:"final_mark_list"`
		ResultFinalMark   string      `json:"result_final_mark,omitempty"`
		Group             group       `json:"group"`
	}
	finalMark struct {
		FinalValue     string `json:"final_value"`
		FinalCriterion string `json:"final_criterion"`
	}
	MarkValue string
	mark      struct {
		MarkValue       MarkValue `json:"mark_value"`
		MarkCriterion   string    `json:"mark_criterion,omitempty"`
		CreatedAt       string    `json:"created_at"`
		FormName        string    `json:"form_name,omitempty"`
		FormDescription string    `json:"form_description,omitempty"`
	}
)

func (mv MarkValue) Int() (int, error) {
	return strconv.Atoi(string(mv))
}
func (mv MarkValue) IsDigit() bool {
	if _, err := strconv.Atoi(string(mv)); err == nil {
		return true
	}
	return false
}

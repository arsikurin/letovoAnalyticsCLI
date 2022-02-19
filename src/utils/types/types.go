package types

import "reflect"

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

type ScheduleResponse struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []data `json:"data"`
}

type data struct {
	PeriodNumDay int         `json:"period_num_day"`
	PeriodName   string      `json:"period_name"`
	PeriodStart  string      `json:"period_start"`
	PeriodEnd    string      `json:"period_end"`
	Date         string      `json:"date"`
	Schedules    []schedules `json:"schedules"`
}

type schedules struct {
	ScheduleRoom int       `json:"schedule_room"`
	Lessons      []lessons `json:"lessons"`
	ZoomMeetings *struct {
		MeetingUrl string `json:"meeting_url"`
	} `json:"zoom_meetings"`
	Group group `json:"group"`
	Room  struct {
		RoomName        string `json:"room_name"`
		RoomDescription string `json:"room_description"`
	} `json:"room"`
}

type lessons struct {
	LessonTopic   *string `json:"lesson_topic,omitempty"`
	LessonCanvas  int     `json:"lesson_canvas"`
	LessonUrl     *string `json:"lesson_url,omitempty"`
	LessonHw      *string `json:"lesson_hw,omitempty"`
	LessonHwDate  *string `json:"lesson_hw_date,omitempty"`
	LessonHwUrl   *string `json:"lesson_hw_url,omitempty"`
	LessonComment *string `json:"lesson_comment,omitempty"`
	Attendance    []struct {
		AttendanceReason *string `json:"attendance_reason,omitempty"`
	} `json:"attendance"`
}

type group struct {
	GroupName string `json:"group_name"`
	Subject   struct {
		SubjectName    string  `json:"subject_name"`
		SubjectNameEng *string `json:"subject_name_eng,omitempty"`
	} `json:"subject"`
}

func (s schedules) IsEmpty() bool {
	return reflect.DeepEqual(s, schedules{})
}

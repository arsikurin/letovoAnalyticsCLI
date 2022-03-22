package types

import "reflect"

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

//{
//    "status": "ok",
//    "code": 200,
//    "message": "Информация по учетной записи получена",
//    "data": {
//        "user": {
//            "id": 671932,
//            "name": "2024antoshin.ii",
//            "email": "2024ant*******@student.letovo.ru",
//            "user_type": "student",
//            "is_email_verified": 1,
//            "is_current_password": 1,
//            "phone": "79*****9716",
//            "parent_id": 0,
//            "student_id": 54142,
//            "is_phone_verified": 0,
//            "bad_pwd_count": 0,
//            "is_blocked_by_badpwd_limit": 0,
//            "bad_ad_pwd_count": 0,
//            "is_disable": 0,
//            "roles": []
//        },
type (
	ScheduleResponse struct {
		Status  string `json:"status"`
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    []data `json:"data"`
	}
	data struct {
		PeriodNumDay int         `json:"period_num_day"`
		PeriodName   string      `json:"period_name"`
		PeriodStart  string      `json:"period_start"`
		PeriodEnd    string      `json:"period_end"`
		Date         string      `json:"date"`
		Schedules    []schedules `json:"schedules"`
	}
	schedules struct {
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
	lessons struct {
		LessonTopic string `json:"lesson_thema,omitempty"`
		//LessonCanvas  int     `json:"lesson_canvas"`
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

func (s schedules) IsEmpty() bool {
	return reflect.DeepEqual(s, schedules{})
}

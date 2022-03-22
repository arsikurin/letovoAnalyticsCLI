package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/arsikurin/letovoAnalyticsCLI/src/utils/database"
	"github.com/arsikurin/letovoAnalyticsCLI/src/utils/types"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func receiveToken(db *sql.DB) (types.TokenResponse, error) {
	creds, err := database.GetCreds(db)
	if err != nil {
		return types.TokenResponse{}, err
	}
	data := url.Values{
		"login":    {*creds.Login},
		"password": {*creds.Password},
	}
	resp, err := http.PostForm("https://s-api.letovo.ru/api/login", data)
	if err != nil {
		return types.TokenResponse{}, err
	}
	if resp.StatusCode != 200 {
		return types.TokenResponse{}, errors.New("unable to obtain token: " + resp.Status)
	}

	res := new(types.TokenResponse)
	err = json.NewDecoder(resp.Body).Decode(res)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Errorln(err)
		}
	}(resp.Body)
	if err != nil {
		return types.TokenResponse{}, err
	}

	return *res, nil
}

func ReceiveStudentID(db *sql.DB) (int, error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://s-api.letovo.ru/api/me", nil)
	if err != nil {
		return 1, err
	}

	tokenResp, err := receiveToken(db)
	if err != nil {
		return 2, err
	}

	req.Header.Set("Authorization", "Bearer "+tokenResp.Data.Token)
	resp, err := client.Do(req)
	if err != nil {
		return 3, err
	}
	if resp.StatusCode != 200 {
		return 4, errors.New("unable to obtain studentID: " + resp.Status)
	}

	res := new(types.MeResponse)
	err = json.NewDecoder(resp.Body).Decode(res)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Errorln(err)
		}
	}(resp.Body)
	if err != nil {
		return 5, err
	}
	return res.Data.User.StudentID, nil
}

func ReceiveScheduleAndHw(db *sql.DB, week bool, specificDay time.Weekday) (types.ScheduleResponse, error) {
	creds, err := database.GetCreds(db)
	if err != nil {
		return types.ScheduleResponse{}, err
	}
	var (
		urlAddr   string
		date      time.Time
		studentId = *creds.StudentID
	)

	if time.Now().Weekday() == 0 {
		date = time.Now().Add(time.Hour * 24)
	} else {
		date = time.Now()
	}

	if week {
		urlAddr = "https://s-api.letovo.ru/api/schedule/" + strconv.Itoa(studentId) + "/week?schedule_date=" + date.Format("2006-01-02")
	} else {
		delta := time.Duration(specificDay - date.Weekday())
		date = date.Add(time.Hour * 24 * delta)
		urlAddr = "https://s-api.letovo.ru/api/schedule/" + strconv.Itoa(studentId) + "/day?schedule_date=" + date.Format("2006-01-02")
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", urlAddr, nil)
	if err != nil {
		return types.ScheduleResponse{}, err
	}

	tokenResp, err := receiveToken(db)
	if err != nil {
		return types.ScheduleResponse{}, err
	}

	req.Header.Set("Authorization", "Bearer "+tokenResp.Data.Token)
	resp, err := client.Do(req)
	if err != nil {
		return types.ScheduleResponse{}, err
	}
	if resp.StatusCode != 200 {
		return types.ScheduleResponse{}, errors.New("unable to obtain schedule: " + resp.Status)
	}

	res := new(types.ScheduleResponse)
	err = json.NewDecoder(resp.Body).Decode(res)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Errorln(err)
		}
	}(resp.Body)
	if err != nil {
		return types.ScheduleResponse{}, err
	}
	return *res, nil
}

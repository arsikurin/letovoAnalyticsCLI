package api

import (
	"encoding/json"
	"github.com/arsikurin/letovoAnalyticsCLI/src/utils/types"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func receiveToken() types.TokenResponse {
	var err error
	data := url.Values{
		"login":    {""},
		"password": {""},
	}
	resp, err := http.PostForm("https://s-api.letovo.ru/api/login", data)
	if err != nil {
		log.Errorln(err)
	}
	if resp.StatusCode != 200 {
		log.Fatalln("unable to obtain token:", resp.Status)
	}

	res := new(types.TokenResponse)
	err = json.NewDecoder(resp.Body).Decode(res)
	if err != nil {
		log.Errorln(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Errorln(err)
		}
	}(resp.Body)
	return *res
}

func ReceiveScheduleAndHw(week bool, specificDay time.Weekday) types.ScheduleResponse {
	var (
		err       error
		urlAddr   string
		date      time.Time
		studentId = 54405
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
		log.Errorln(err)
	}

	req.Header.Set("Authorization", "Bearer "+receiveToken().Data.Token)
	resp, err := client.Do(req)
	if err != nil {
		log.Errorln(err)
	}
	if resp.StatusCode != 200 {
		log.Fatalln("unable to obtain schedule:", resp.Status)
	}
	//var v interface{}
	//body, err := ioutil.ReadAll(resp.Body)
	//json.Unmarshal(body, &v)
	//fmt.Println(v)
	res := new(types.ScheduleResponse)
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		log.Errorln(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Errorln(err)
		}
	}(resp.Body)
	//jsonData, err := json.MarshalIndent(res, "", "    ")
	//fmt.Println(string(jsonData))
	return *res
}

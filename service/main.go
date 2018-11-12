package service

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/hal-ms/game/log"
	"github.com/hal-ms/game/repo"
)

var Main = mainService{}

type mainService struct {
}

type StartMsg struct {
	Success string `json:"success"`
}

type CheckMsg struct {
}

type EndMsg struct {
	Job string `json:"job"`
}

func (m *mainService) Start() bool {
	res, _ := m.req("GET", "http://example.com:8080/start", nil)
	b, _ := ioutil.ReadAll(res.Body)

	var msg StartMsg
	_ = json.Unmarshal(b, &msg)

	return msg.Success == "true"
}

func (m *mainService) Check() {
	//res, _ := m.req("GET", "http://example.com:8080/check", nil)

}

func (m *mainService) End() {
	res, _ := m.req("GET", "http://example.com:8080/end", nil)
	b, _ := ioutil.ReadAll(res.Body)

	var msg EndMsg
	_ = json.Unmarshal(b, &msg)

	job := msg.Job

	if _, err := repo.Job.Exist(job); err != nil {
		log.SendSlack(err.Error())
		return
	}

	repo.Job.Job(job)
}

func (m *mainService) req(method, url string, body io.Reader) (*http.Response, error) {
	req, _ := http.NewRequest(method, url, body)

	client := new(http.Client)
	res, err := client.Do(req)
	defer res.Body.Close()

	return res, err
}

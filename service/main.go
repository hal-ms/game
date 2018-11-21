package service

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/hal-ms/game/log"
	"github.com/hal-ms/game/repo"
	"github.com/makki0205/config"
)

var Main = mainService{}

type mainService struct {
}

type StartMsg struct {
	Job string `json:"job"`
	//Success bool   `json:"success"`
}

type CheckMsg struct {
}

type EndMsg struct {
	Job string `json:"job"`
}

func (m *mainService) Start() bool {
	res, _ := m.req("GET", config.Env("mainUrl")+"/start", nil)
	b, _ := ioutil.ReadAll(res.Body)

	var msg StartMsg
	_ = json.Unmarshal(b, &msg)

	job := msg.Job

	repo.Job.Job(job) // 次の仕事をセット

	if _, err := repo.Job.Exist(job); err != nil {
		log.SendSlack(err.Error())
		return false
	}

	return true
}

func (m *mainService) Check() {
	//res, _ := m.req("GET", config.Env("mainUrl")+"/check", nil)

}

func (m *mainService) End() {
	res, _ := m.req("GET", config.Env("mainUrl")+"/start", nil)
	b, _ := ioutil.ReadAll(res.Body)

	var msg EndMsg
	_ = json.Unmarshal(b, &msg)

	// 終了処理
	repo.State.IsStandby(true) // 待機状態に遷移
	repo.Hit.Reset()           // ヒットポイントをリセット

}

func (m *mainService) req(method, url string, body io.Reader) (*http.Response, error) {
	req, _ := http.NewRequest(method, url, body)

	client := new(http.Client)
	res, err := client.Do(req)
	defer res.Body.Close()

	return res, err
}

package service

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/hal-ms/game/log"
	"github.com/hal-ms/game/repo"
	"github.com/makki0205/config"
)

var Main = mainService{}

type mainService struct {
}

func init() {

}

type StartMsg struct {
	Name string `json:"name"`
	//Success bool   `json:"success"`
}

type CheckMsg struct {
}

type EndMsg struct {
	Job string `json:"job"`
}

func (m *mainService) Start() bool {
	res, _ := m.req("GET", config.Env("mainUrl")+"/api/game/start", nil)
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	var msg StartMsg
	err = json.Unmarshal(b, &msg)
	if err != nil {
		panic(err)
	}
	fmt.Println(msg.Name)
	job := msg.Name

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
	m.req("GET", config.Env("mainUrl")+"/api/game/end", nil)

	// 終了処理
	repo.State.IsStandby(true) // 待機状態に遷移
	repo.Hit.Reset()           // ヒットポイントをリセット
	//中央画面処理待ち
	time.Sleep(5 * time.Second)
	LCD.Reset()

}

func (m *mainService) req(method, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		panic(err)
	}

	fmt.Println(req)
	client := new(http.Client)
	res, err := client.Do(req)
	//defer res.Body.Close()

	return res, err
}

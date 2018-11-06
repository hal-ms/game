package service

import (
	"errors"
	"fmt"
)

var LCD = newLcdService()

type lcdService struct {
	state    int
	progress int
}

const (
	Standby = iota
	Start
	Stop
	Hide
)

func newLcdService() lcdService {
	return lcdService{}
}

// jobの登録
func (l *lcdService) SetJob() error {
	fmt.Println("setJob")
	return nil
}

// アニメーションスタート
func (l *lcdService) Start() error {
	l.state = Start
	fmt.Println("start")
	return nil
}

// アニメーションストップ
func (l *lcdService) Stop() error {
	l.state = Stop
	fmt.Println("stop")
	return nil
}

// ディスプレイを消す
func (l *lcdService) Hide() error {
	l.state = Hide
	fmt.Println("hide")
	return nil
}

// 次のアニメーション
func (l *lcdService) Next(progress int) error {
	if l.progress > progress || l.progress+1 < progress {
		return errors.New("LCD Next 要求が不正です")
	}
	if l.progress+1 == progress {
		l.progress++
		fmt.Println("next")
		// nextを送信
	}
	return nil
}

// リセット
func (l *lcdService) Reset() error {
	l.progress = 0
	fmt.Println("reset")
	l.state = Standby
	return nil
}

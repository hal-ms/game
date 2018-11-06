package service

import "errors"

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

	return nil
}

// アニメーションスタート
func (l *lcdService) Start() error {
	l.state = Start

	return nil
}

// アニメーションストップ
func (l *lcdService) Stop() error {
	l.state = Stop

	return nil
}

// ディスプレイを消す
func (l *lcdService) Hide() error {
	l.state = Hide
	return nil
}

// 次のアニメーション
func (l *lcdService) Next(progress int) error {
	if l.progress > progress || l.progress+1 < progress {
		return errors.New("LCD Next 要求が不正です")
	}
	if l.progress+1 == progress {
		l.progress++
		// nextを送信
	}
	return nil
}

// リセット
func (l *lcdService) Reset() error {
	l.progress = 0
	l.state = Standby
	return nil
}

package service

import (
	"errors"
	"fmt"

	"github.com/hal-ms/game/log"
	"github.com/tarm/serial"
)

var LCD = newLcdService()

type lcdService struct {
	state    int
	progress int
	conn     *serial.Port
}

const (
	Standby = iota
	Start
	Stop
	Hide
	Show
)

func newLcdService() lcdService {
	c := &serial.Config{Name: "COM9", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.SendSlack(err.Error())
		panic(err)
	}
	conn := s
	return lcdService{conn: conn}
}

// jobの登録
func (l *lcdService) SetJob(j string) error {
	fmt.Println("setJob")
	var job byte = 0x41
	switch j {
	case "cook":
		job = 0x41
	default:
		job = 0x00
	}

	if job != 0x00 {
		l.write([]byte{job})
	}

	return nil
}

// アニメーションスタート
func (l *lcdService) Start() error {
	l.state = Start
	fmt.Println("start")
	l.write([]byte{0x31})
	return nil
}

// アニメーションストップ
func (l *lcdService) Stop() error {
	l.state = Stop
	fmt.Println("stop")
	l.write([]byte{0x30})
	return nil
}

// ディスプレイを消す
func (l *lcdService) Hide() error {
	l.state = Hide
	fmt.Println("hide")
	l.write([]byte{0x2D})
	return nil
}

// ディスプレイを表示する
func (l *lcdService) Show() error {
	l.state = Show
	fmt.Println("show")
	l.write([]byte{0x2B})
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
		l.write([]byte{0x20})
	}
	return nil
}

// リセット
func (l *lcdService) Reset() error {
	l.progress = 0
	fmt.Println("reset")
	l.write([]byte{0x1B})
	l.state = Standby
	return nil
}

// 信号送信
func (l *lcdService) write(b []byte) error {
	_, err := l.conn.Write(b)
	if err != nil {
		log.SendSlack(err.Error())
		panic(err)
	}
	return nil
}

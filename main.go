package main

import (
	"strconv"

	"github.com/makki0205/log"

	"github.com/hal-ms/game/service"

	"github.com/makki0205/config"

	"github.com/gin-gonic/gin"
	"github.com/hal-ms/game/cnto"
	"github.com/tarm/serial"
)

func main() {
	log.ServiceName = "ms-game"
	service.LCD.Reset()
	r := gin.Default()
	r.POST("/button", cnto.Button)
	r.POST("/is_wearing/:IsWearing", cnto.IsWearing)
	r.POST("/job/:job", cnto.Job)
	go r.Run()
	hitScreen()
}

func hitScreen() {
	c := &serial.Config{Name: config.Env("micPort"), Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		panic(err)
	}
	buf := make([]byte, 4)
	if err != nil {
		log.SendSlack(err.Error())
		panic(err)
	}
	for {
		n, err := s.Read(buf)
		if err != nil {
			panic(err)
		}
		p, err := strconv.Atoi(string(buf[:n]))
		if err != nil {
			log.SendSlack(err.Error())
			continue
		}
		cnto.Hit(p)
	}
}
